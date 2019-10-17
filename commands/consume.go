package commands

import (
	"strings"

	"github.com/davecgh/go-spew/spew"

	kafka "github.com/minhajuddinkhan/kafka-client"
	"github.com/urfave/cli"
)

//Consume consumes messages on kafka broker
func Consume() cli.Command {
	var topic, brokers string
	var partition, offset int

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
			cli.IntFlag{
				Name:        "offset",
				Destination: &offset,
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
			if partition == -1 {
				return cli.NewExitError("invalid flag empty", 1)
			}
			if offset == -1 {
				return cli.NewExitError("brokers brokers empty", 1)
			}
			return nil
		},
		Action: func(c *cli.Context) error {
			kc, err := kafka.NewClient(strings.Split(brokers, ","))
			if err != nil {
				return err
			}
			msgCh := make(chan interface{})
			go func() {
				spew.Dump(<-msgCh)
			}()
			err = kc.Consume(topic, int32(partition), msgCh)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
