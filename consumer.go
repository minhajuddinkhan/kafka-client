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
		fmt.Println(topic, partition)
		cp, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			return err
		}
		select {
		case msg := <-cp.Messages():
			cp.Close()
			spew.Dump(msg)
		case err := <-cp.Errors():
			return err
		}
	}

}
