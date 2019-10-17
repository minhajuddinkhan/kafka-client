package kafka

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"

	"github.com/Shopify/sarama"
)

func (kc *client) Consume(topic string, partition int32, msgCh chan interface{}) (err error) {

	consumer, err := sarama.NewConsumer(kc.brokers, kc.config)
	if err != nil {
		return fmt.Errorf("unable to create consumer. err: %v", err)
	}

	for {
		cp, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			spew.Dump(err)
		}
		select {

		case msg := <-cp.Messages():
			msgCh <- msg.Value
			cp.Close()
		case err := <-cp.Errors():
			spew.Dump(err)
		}
	}

}
