package org

import (
	"fmt"

	"github.com/openworklabs/streams-cli/v2/types"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/urfave/cli/v2"
)

func Create(ctx *cli.Context, tclient *client.Client) error {
	id, err := thread.Decode("bafk2pukjgfvfgantvjqk7ggtv7h2brji2iw74ztfcfsq5so6kx6alkq")
	// fmt.Println("THREAD", threadId)
	orgName := ctx.Args().First()

	ids, err := tclient.Create(
		ctx.Context,
		id,
		"Organization",
		client.Instances{createOwner(id, orgName)},
	)

	if err != nil {
		return err
	}

	fmt.Println(ids)
	return nil
}

func createOwner(threadID thread.ID, name string) *types.Owner {
	return &types.Owner{
		ThreadID: threadID,
	}
}
