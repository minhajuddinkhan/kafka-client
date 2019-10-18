package commands

import (
	"github.com/davecgh/go-spew/spew"

	kafka "github.com/minhajuddinkhan/kafka-client"
	"github.com/minhajuddinkhan/kafka-client/store"
	"github.com/urfave/cli"
)

//Produce produces messages on kafka broker
func Produce(store store.Kafka) cli.Command {
	var topic string
	var value string
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
		},
		Before: func(c *cli.Context) error {
			if topic == "" {
				return cli.NewExitError("topic flag empty", 1)
			}
			if value == "" {
				return cli.NewExitError("value flag empty", 1)
			}

			return nil
		},
		Action: func(c *cli.Context) error {

			brokers, err := store.GetBrokers()
			if err != nil {
				return err
			}

			urls := make([]string, len(brokers))
			for i, x := range brokers {
				urls[i] = x.URL
			}
			kc := kafka.NewClient(urls)
			spew.Dump(kc.Produce(topic, value))
			return nil
		},
	}
}
