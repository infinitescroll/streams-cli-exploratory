package stream

import (
	"context"

	"github.com/multiformats/go-multiaddr"
	"github.com/openworklabs/streams-cli/v2/types"
	"github.com/openworklabs/streams-cli/v2/utils"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/go-threads/util"
	pow "github.com/textileio/powergate/api/client"
	"github.com/textileio/powergate/ffs/api"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func Create(ctx *cli.Context, tclient *client.Client) error {
	orgThreadID, err := fetchOrgThread(ctx, tclient)
	if err != nil {
		return err
	}
	ffsId, authToken, walletAddress, err := createFFSInstance(ctx)
	if err != nil {
		return err
	}

	streamThreadId := thread.NewIDV1(thread.Raw, 32)
	tclient.NewDB(ctx.Context, streamThreadId)
	tclient.NewCollection(ctx.Context, streamThreadId, db.CollectionConfig{
		Name:   "StreamMeta",
		Schema: util.SchemaFromInstance(&types.StreamMeta{}, false),
	})

	_, err = tclient.Create(
		ctx.Context,
		orgThreadID,
		"StreamPointer",
		client.Instances{createStreamPointer(
			streamThreadId,
			ctx.Args().Get(0),
		)},
	)
	if err != nil {
		return err
	}

	_, err = tclient.Create(
		ctx.Context,
		streamThreadId,
		"StreamMeta",
		client.Instances{createStreamMeta(
			ctx.Args().Get(0),
			ffsId,
			authToken,
			walletAddress,
		)},
	)
	return nil
}

func fetchOrgThread(ctx *cli.Context, tclient *client.Client) (thread.ID, error) {
	owner := ctx.Args().Get(1)
	id := utils.GetMetaThread()
	q := db.Where("name").Eq(owner)

	results, err := tclient.Find(
		ctx.Context,
		id,
		"Organization",
		q,
		&types.OwnerPointer{},
	)

	if err != nil {
		return thread.ID{}, err
	}

	orgs := results.([]*types.OwnerPointer)
	threadID, err := thread.Decode(orgs[0].ThreadID)

	if err != nil {
		return thread.ID{}, err
	}

	return threadID, nil
}

func createFFSInstance(ctx *cli.Context) (string, string, []api.AddrInfo, error) {
	ma, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/5002")
	if err != nil {
		return "", "", []api.AddrInfo{}, err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	auth := pow.TokenAuth{}
	opts = append(opts, grpc.WithPerRPCCredentials(auth))
	pclient, err := pow.NewClient(ma, opts...)
	if err != nil {
		return "", "", []api.AddrInfo{}, err
	}

	ffsID, token, err := pclient.FFS.Create(ctx.Context)
	if err != nil {
		return "", "", []api.AddrInfo{}, err
	}

	addrInfo, err := pclient.FFS.Addrs(context.WithValue(ctx.Context, pow.AuthKey, token))
	if err != nil {
		return "", "", []api.AddrInfo{}, err
	}

	return ffsID, token, addrInfo, nil
}

func createStreamPointer(threadID thread.ID, name string) *types.StreamPointer {
	return &types.StreamPointer{
		ID:       "",
		ThreadID: threadID.String(),
		Name:     name,
	}
}

func createStreamMeta(name string, ffsID string, ffsAuthToken string, walletAddresses []api.AddrInfo) *types.StreamMeta {
	return &types.StreamMeta{
		ID:              "",
		Name:            name,
		FFSID:           ffsID,
		FFSAuthToken:    ffsAuthToken,
		WalletAddresses: walletAddresses,
	}
}
