package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:     "zookeeper",
			Category: "zookeeper",
			Aliases:  []string{"zk"},
			Usage:    "zookeeper service",
			Subcommands: []cli.Command{
				{
					Name: "stat",
					Action: func(c *cli.Context) error {
						fmt.Println("stat task: ", c.Args())
						return nil
					},
				},
				{
					Name: "conf",
					Action: func(c *cli.Context) error {
						fmt.Println("conf task: ", c.Args())
						return nil
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
