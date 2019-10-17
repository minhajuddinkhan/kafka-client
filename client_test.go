package kafka_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	kafka "github.com/minhajuddinkhan/kafka-client"
)

func TestNewKafkaClient(t *testing.T) {

	brokerUrls := []string{}
	client := kafka.NewClient(brokerUrls)
	assert.NotNil(t, client)
}
