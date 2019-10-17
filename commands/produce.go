package commands

import (
	"strings"

	"github.com/davecgh/go-spew/spew"

	kafka "github.com/minhajuddinkhan/kafka-client"
	"github.com/urfave/cli"
)

//Produce produces messages on kafka broker
func Produce() cli.Command {
	var topic string
	var value string
	var brokers string
	return cli.Command{
		Name:        "produce",
		Description: "publishes given message on given topic",
		Usage:       "publishes given message on given topic",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "topic",
				Destination: &topic,
			},
			cli.StringFlag{
				Name:        "value",
				Destination: &value,
			},
			cli.StringFlag{
				Name:        "brokers",
				Destination: &brokers,
			},
		},
		Before: func(c *cli.Context) error {
			if topic == "" {
				return cli.NewExitError("topic flag empty", 1)
			}
			if value == "" {
				return cli.NewExitError("value flag empty", 1)
			}
			if brokers == "" {
				return cli.NewExitError("brokers brokers empty", 1)
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			kc := kafka.NewClient(strings.Split(brokers, ","))
			spew.Dump(kc.Produce(topic, value))
			return nil
		},
	}
}
