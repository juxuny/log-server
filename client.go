package log_server

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"time"
)

type client struct {
	host string
	cert string
	LogServerClient
}

func NewClient(host string, cert string) (c *client, err error) {
	c = &client{
		host: host,
		cert: cert,
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	c.LogServerClient, err = getClient(ctx, host, cert)
	return c, err
}

func loadTLSCredentials(certFile string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}

func getClient(ctx context.Context, host string, certFile string) (client LogServerClient, err error) {
	tlsCredentials, err := loadTLSCredentials(certFile)
	if err != nil {
		return client, errors.Wrap(err, "load cert failed")
	}

	conn, err := grpc.DialContext(ctx, host, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		return nil, errors.Wrap(err, "connect failed")
	}
	client = NewLogServerClient(conn)
	return client, nil
}
