package kafka_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	kafka "github.com/minhajuddinkhan/kafka-client"
)

func TestNewKafkaClient(t *testing.T) {

	brokerUrls := []string{}
	client, err := kafka.NewClient(brokerUrls)
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestClientHasBrokerUrls(t *testing.T) {

	brokerUrls := []string{"3.130.146.170:9092"}
	_, err := kafka.NewClient(brokerUrls)
	assert.Nil(t, err)
}
