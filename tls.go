package kafka

import (
	"crypto/tls"

	"github.com/certifi/gocertifi"
)

//SetTlsConfig SetTlsConfig
func (kc *client) SetTLSConfig(certFilePath, keyFilePath string) error {

	cert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return err
	}
	certPool, err := gocertifi.CACerts()
	if err != nil {
		return err
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	kc.config.Net.TLS.Config = tlsConfig
	kc.config.Net.TLS.Enable = true

	return nil
}
