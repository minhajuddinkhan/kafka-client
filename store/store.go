package store

import (
	"github.com/minhajuddinkhan/kafka-client/entities"
)

//Kafka Store
type Kafka interface {
	SetTLS(certFile, keyFile string) error
	SetBrokers(brokers []string) error
	GetBrokers() ([]entities.Broker, error)
}
