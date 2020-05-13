package main

import (
	"context"
	"log"
	"os"

	"github.com/openworklabs/streams-cli/v2/cmds/org"
	"github.com/openworklabs/streams-cli/v2/cmds/stream"
	"github.com/openworklabs/streams-cli/v2/utils"
	"github.com/textileio/go-threads/api/client"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	tclient, err := client.NewClient("0.0.0.0:6006", grpc.WithInsecure())
	utils.CheckErr(err)
	defer tclient.Close()
	utils.CreateStreamsMetaThread(tclient)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "org",
				Subcommands: []*cli.Command{
					{
						Name: "create",
						Action: func(c *cli.Context) error {
							return org.Create(c, tclient)
						},
						ArgsUsage: "<email>",
					},
					{
						Name: "get",
						Action: func(c *cli.Context) error {
							return org.Get(c, tclient)
						},
						ArgsUsage: "<name>",
					},
				},
			},
			{
				Name: "stream",
				Subcommands: []*cli.Command{
					{
						Name: "create",
						Action: func(c *cli.Context) error {
							return stream.Create(c, tclient)
						},
						ArgsUsage: "<stream-name> <owner-name> <owner-type>",
					},
					{
						Name: "get",
						Action: func(c *cli.Context) error {
							return org.Get(c, tclient)
						},
						ArgsUsage: "<name>",
					},
				},
			},
			{
				Name: "reset",
				Action: func(c *cli.Context) error {
					id := utils.GetMetaThread()
					err := tclient.DeleteDB(context.Background(), id)
					if err != nil {
						return err
					}
					return nil
				},
			},
		},
	}

	apperr := app.Run(os.Args)
	if apperr != nil {
		log.Fatal(apperr)
	}
}
