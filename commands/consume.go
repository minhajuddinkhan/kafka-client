package commands

import (
	"github.com/davecgh/go-spew/spew"

	kafka "github.com/minhajuddinkhan/kafka-client"
	"github.com/minhajuddinkhan/kafka-client/entities"
	"github.com/minhajuddinkhan/kafka-client/store"
	"github.com/urfave/cli"
)

//Consume consumes messages on kafka broker
func Consume(store store.Kafka) cli.Command {
	var topic string
	var partition int

	return cli.Command{
		Name:        "consume",
		Description: "consumes a message on given topic, partition and offset",
		Usage:       "consumes a message on given topic, partition and offset",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "topic",
				Destination: &topic,
			},
			cli.IntFlag{
				Name:        "partition",
				Destination: &partition,
			},
		},
		Before: func(c *cli.Context) error {

			if topic == "" {
				return cli.NewExitError("topic flag empty", 1)
			}
			if partition == -1 {
				return cli.NewExitError("invalid flag empty", 1)
			}
			return nil
		},
		Action: func(c *cli.Context) error {

			brokers, err := store.GetBrokers()
			if err != nil {
				return err
			}
			spew.Dump("ORIGINAL INDB", brokers)
			bg := entities.BrokerGroup{Brokers: brokers}
			kc := kafka.NewClient(bg.URLs())

			msgCh := make(chan interface{})

			// TODO:: add mechanism to receieve messages
			err = kc.Consume(topic, int32(partition), msgCh)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
