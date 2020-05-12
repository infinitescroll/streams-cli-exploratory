package org

import (
	"fmt"

	"github.com/openworklabs/streams-cli/v2/types"
	"github.com/openworklabs/streams-cli/v2/utils"
	"github.com/textileio/go-threads/api/client"
	"github.com/textileio/go-threads/core/thread"
	"github.com/textileio/go-threads/db"
	"github.com/urfave/cli/v2"
)

func Get(ctx *cli.Context, tclient *client.Client) error {
	id := utils.GetMetaThread()

	results, err := tclient.Find(
		ctx.Context,
		id,
		"Organization",
		&db.Query{},
		&types.Owner{},
	)
	if err != nil {
		return err
	}

	orgs := results.([]*types.Owner)

	for _, v := range orgs {
		id, err := thread.Decode(v.ThreadID)
		if err != nil {
			return err
		}

		results, err := tclient.Find(
			ctx.Context,
			id,
			"OwnerMetadata",
			&db.Query{},
			&types.OwnerMetadata{},
		)
		if err != nil {
			return err
		}

		metadata := results.([]*types.OwnerMetadata)
		fmt.Println(metadata[0].Name)
	}

	return nil
}
