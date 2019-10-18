package kafka

import (
	"crypto/tls"

	"github.com/Shopify/sarama"
)

//Client the kafka client
type Client interface {
	// Gets the list of brokers
	Brokers() []string

	// Publishes a message on the given topic
	// Returns partition, offset and error occurred if any
	Produce(topic string, value interface{}) (partition int32, offset int64, err error)

	// Consumes a message on the given topic and partition
	// Returns error occurred if any
	Consume(topic string, partition int32, msgCh chan interface{}) (error error)

	SetTLSConfig(certFilepath, keyFilepath string) error
}

type client struct {
	brokers   []string
	tlsConfig *tls.Config
	config    *sarama.Config
}

// NewClient creates a new client
func NewClient(brokerUrls []string) Client {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true
	conf.Consumer.Return.Errors = true
	conf.Consumer.Offsets.Initial = sarama.OffsetOldest

	return &client{brokers: brokerUrls, config: conf}
}
