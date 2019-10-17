package kafka

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
)

func (kc *client) Produce(topic string, value interface{}) (int32, int64, error) {

	producer, err := sarama.NewSyncProducer(kc.brokers, kc.config)
	if err != nil {
		return -1, -1, err
	}
	defer producer.Close()
	b, err := json.Marshal(value)
	if err != nil {
		return -1, -1, fmt.Errorf("Unable to convert value to bytes: %v", err)
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(b),
	}
	return producer.SendMessage(msg)
}
