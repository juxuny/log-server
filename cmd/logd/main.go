package main

import (
	"flag"
	"fmt"
	log_server "github.com/juxuny/log-server"
	"github.com/juxuny/log-server/log"
	"google.golang.org/grpc"
	"net"
	"os"
)

var (
	port     int
	certFile string
	keyFile  string
	logger   = log.NewLogger("[logd]")
)

func init() {
	flag.StringVar(&certFile, "cert", "cert/server-cert.pem", "cert file")
	flag.StringVar(&keyFile, "key", "cert/server-key.pem", "cert key file")
	flag.IntVar(&port, "p", 40000, "listen port")
}

func main() {
	flag.Parse()
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to listen:", err)
	}
	defer func() {
		if err := ln.Close(); err != nil {
			logger.Error(err)
		}
	}()
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		logger.Error("cannot load TLS credentials: ", err)
		os.Exit(-1)
	}
	s := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)
	fmt.Println("listen", addr)
	log_server.RegisterLogServerServer(s, &server{})
	if err := s.Serve(ln); err != nil {
		logger.Error("failed to serve:", err)
	}
}
