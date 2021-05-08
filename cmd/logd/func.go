package main

import (
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc/credentials"
	"os"
)

func touchDir(dir string) error {
	stat, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.MkdirAll(dir, 0776)
	}
	if stat.IsDir() {
		return nil
	} else {
		return fmt.Errorf("path %s is not a director", dir)
	}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	fmt.Println("cert file:", certFile)
	fmt.Println("cert key file:", keyFile)
	serverCert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}
