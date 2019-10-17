package main

import (
	"log"
	"os"

	"github.com/minhajuddinkhan/kafka-client/commands"
	"github.com/urfave/cli"
)

var brokerUrls string

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "brokers",
			Destination: &brokerUrls,
			Usage:       "Saves broker urls",
		},
	}
	app.Commands = []cli.Command{
		commands.Consume(),
		commands.Produce(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
