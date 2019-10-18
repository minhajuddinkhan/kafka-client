package main

import (
	"log"
	"os"

	"github.com/minhajuddinkhan/kafka-client/commands"
	"github.com/minhajuddinkhan/kafka-client/store/sqlite"
	"github.com/urfave/cli"
)

var brokerUrls string

func main() {

	kafkaStore, err := sqlite.NewStore()
	if err != nil {
		log.Fatal(err)
	}

	app := cli.NewApp()
	app.Name = "Apache Kafka Client CLI"
	app.Usage = "CLI application for interacting with Apache Kafka"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "brokers",
			Destination: &brokerUrls,
			Usage:       "Saves broker urls",
		},
	}
	app.Commands = []cli.Command{
		commands.Consume(kafkaStore),
		commands.Produce(kafkaStore),
		commands.Brokers(kafkaStore),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
