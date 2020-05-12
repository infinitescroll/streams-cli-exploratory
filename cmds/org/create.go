package org

import (
	"github.com/openworklabs/streams-cli/v2/types"
	"github.com/openworklabs/streams-cli/v2/utils"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/textileio/go-threads/util"
	"github.com/urfave/cli/v2"
)

func Create(ctx *cli.Context, tclient *client.Client) error {
	id := utils.GetMetaThread()
	orgThreadId := thread.NewIDV1(thread.Raw, 32)
	orgName := ctx.Args().First()

	_, err := tclient.Create(
		ctx.Context,
		id,
		"Organization",
		client.Instances{createOwnerPointer(orgThreadId, orgName)},
	)
	if err != nil {
		return err
	}

	err = tclient.NewDB(ctx.Context, orgThreadId)
	if err != nil {
		return err
	}

	err = tclient.NewCollection(ctx.Context, orgThreadId, db.CollectionConfig{
		Name:   "OwnerMetadata",
		Schema: util.SchemaFromInstance(&types.OwnerMetadata{}, false),
	})
	if err != nil {
		return err
	}

	_, err = tclient.Create(
		ctx.Context,
		orgThreadId,
		"OwnerMetadata",
		client.Instances{createOwnerMetadata(orgName)},
	)
	if err != nil {
		return err
	}

	return nil
}

func createOwnerPointer(threadID thread.ID, orgName string) *types.OwnerPointer {
	return &types.OwnerPointer{
		ID:       "",
		ThreadID: threadID.String(),
		Name:     orgName,
	}
}

func createOwnerMetadata(orgName string) *types.OwnerMetadata {
	return &types.OwnerMetadata{
		ID:    "",
		Name:  orgName,
		Email: "",
	}
}
