package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/davecgh/go-spew/spew"
)

func (kc *client) Consume(topic string, partition int32, msgCh chan interface{}) (err error) {

	consumer, err := sarama.NewConsumer(kc.brokers, kc.config)
	if err != nil {
		return fmt.Errorf("unable to create consumer. err: %v", err)
	}

	defer consumer.Close()

	for {
		cp, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		select {

		case msg := <-cp.Messages():
			msgCh <- msg.Value
			spew.Dump("closing!")
			cp.Close()
		case err := <-cp.Errors():
			return err
		}
	}

}
