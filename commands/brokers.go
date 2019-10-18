package commands

import (
	"strings"

	"github.com/davecgh/go-spew/spew"

	"github.com/minhajuddinkhan/kafka-client/store"
	"github.com/urfave/cli"
)

//Brokers handles brokers
func Brokers(store store.Kafka) cli.Command {
	var brokers string

	return cli.Command{
		Name:        "brokers",
		Description: "saves broker urls to storage",
		Usage:       "saves broker urls to storage",
		Subcommands: []cli.Command{
			cli.Command{
				Name: "set",
				Before: func(c *cli.Context) error {
					brokers = c.Args().Get(0)
					if brokers == "" {
						return cli.NewExitError("brokers cant be empty", 1)
					}
					return nil
				},
				Action: func(c *cli.Context) error {
					// spew.Dump(brokers)
					// return nil
					return store.SetBrokers(strings.Split(brokers, ","))
				},
			},
			cli.Command{
				Name: "get",
				Action: func(c *cli.Context) error {

					brokers, err := store.GetBrokers()
					if err != nil {
						cli.NewExitError(err.Error(), 1)
					}
					spew.Dump(brokers)
					return nil
				},
			},
		},
	}
}
