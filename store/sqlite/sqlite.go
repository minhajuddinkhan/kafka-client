package sqlite

import (
	"crypto/tls"
	"os"

	"github.com/jinzhu/gorm"
	//SQLITE dialect support
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/minhajuddinkhan/kafka-client/entities"
	"github.com/minhajuddinkhan/kafka-client/store"
	"github.com/minhajuddinkhan/kafka-client/store/models"
)

type kafkastore struct {
	Conn *gorm.DB
}

//NewStore NewStore
func NewStore() (store.Kafka, error) {

	conn, err := gorm.Open("sqlite3", "/tmp/kafka.db")
	if err != nil {
		return nil, err
	}
	models := models.All()
	for _, m := range models {
		if !conn.HasTable(m) {
			conn.AutoMigrate(m)
		}
	}
	return &kafkastore{Conn: conn}, err
}

func (s *kafkastore) SetBrokers(brokers []string) error {

	for _, b := range brokers {
		broker := models.Broker{URL: b}
		err := s.Conn.Create(&broker).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *kafkastore) SetTLS(certFile, keyFile string) error {

	if _, err := os.Stat(certFile); err != nil {
		return err
	}

	if _, err := os.Stat(keyFile); err != nil {
		return err
	}
	if err := validateTLS(certFile, keyFile); err != nil {
		return err
	}
	var tls models.TLS
	err := s.Conn.First(&tls).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return s.Conn.Create(models.TLS{CertPath: certFile, KeyPath: keyFile}).Error
		}
		return err
	}
	tls.CertPath = certFile
	tls.KeyPath = keyFile
	return s.Conn.Update(&tls).Error
}

func (s *kafkastore) GetBrokers() ([]entities.Broker, error) {

	var brokers []models.Broker
	if err := s.Conn.Find(&brokers).Error; err != nil {
		return nil, err
	}
	result := make([]entities.Broker, len(brokers))

	for i, br := range brokers {
		result[i].URL = br.URL
		result[i].ID = br.ID
	}
	return result, nil
}

func validateTLS(certFile, keyFile string) error {

	_, err := tls.LoadX509KeyPair(certFile, keyFile)
	return err

}
