package main

import (
	"log"
	"os"

	"github.com/openworklabs/streams-cli/v2/cmds/org"
	"github.com/openworklabs/streams-cli/v2/utils"
	"github.com/textileio/go-threads/api/client"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	var tclient *client.Client
	var err error

	tclient, err = client.NewClient("0.0.0.0:6006", grpc.WithInsecure())
	utils.CheckErr(err)
	defer tclient.Close()

	utils.CreateStreamsMetaThread(tclient)

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "streams",
				Usage: "interact with streams",
				Subcommands: []*cli.Command{
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
						},
					},
				},
			},
		},
	}

	apperr := app.Run(os.Args)
	if apperr != nil {
		log.Fatal(apperr)
	}
}
