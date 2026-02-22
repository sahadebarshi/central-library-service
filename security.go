package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
)

// 🔐 Setup TLS configuration with mutual TLS
func createTLSConfig(server_cert, server_key, server_ca_cert string) *tls.Config {
	// Load server certificate and private key
	cert, err := tls.LoadX509KeyPair(server_cert, server_key)
	if err != nil {
		log.Fatalf("Failed to load server certificate/key: %v", err)
	}

	// Load the CA certificate (used to verify client certs)
	caCert, err := os.ReadFile(server_ca_cert)
	if err != nil {
		log.Fatalf("Failed to read CA certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
		log.Fatal("Failed to append CA cert to pool")
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // Enforce mTLS
		ClientCAs:    caCertPool,
		MinVersion:   tls.VersionTLS12,
	}
}
