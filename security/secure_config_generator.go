package security

import (
	"crypto/tls"
	"crypto/x509"
	"log"
)

func NewSecureConfig(key string, pubKey string, server bool) *tls.Config {
	cert, err := tls.LoadX509KeyPair(pubKey, key)
	if err != nil {
		log.Fatalf("client: loadkeys: %s", err)
	}

	if len(cert.Certificate) != 2 {
		log.Fatal("client.crt should have 2 concatenated certificates: client + CA")
	}

	ca, err := x509.ParseCertificate(cert.Certificate[1])
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	certPool.AddCert(ca)

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	if server {
		config.RootCAs = certPool
	} else {
		config.ClientAuth = tls.RequireAndVerifyClientCert
		config.ClientCAs = certPool
	}

	return &config
}
