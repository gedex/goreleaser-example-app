package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var (
	version = "master"
	commit  = "none"
	date    = "unknown"
)

func main() {
	app := cli.NewApp()
	app.Name = "goreleaser-example-app"
	app.Usage = "Example app to play with goreleaser"
	app.Version = fmt.Sprintf("%v, commit %v, built at %v", version, commit, date)
	app.Commands = []cli.Command{
		{
			Name:  "hello",
			Usage: "Say hello world!",
			Action: func(c *cli.Context) error {
				fmt.Println("hello world!")
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
	}
}
