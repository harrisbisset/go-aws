package main

import (
	"context"
	"fmt"
	"os"

	"github.com/harrisbisset/go-aws/pkg/dev"
	"github.com/urfave/cli/v3"
)

func main() {
	// variables for commands
	var cidr string
	var addr string

	cmd := &cli.Command{
		Name:  "Go Aws Wrapper",
		Usage: "Manage small aws applications",
		Commands: []*cli.Command{
			{
				Name:  "local",
				Usage: "start localhost website to manage the application",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "addr",
						Value:       "80",
						Usage:       "address for dev server",
						Destination: &addr,
					},
					&cli.StringFlag{
						Name:        "cidr",
						Value:       "0.0.0.0",
						Usage:       "cidr block for dev server",
						Destination: &cidr,
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					dev.StartDev(addr, cidr)
					return nil
				},
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
