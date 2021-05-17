package main

import (
	"context"
	"flag"
	"fmt"
	log_server "github.com/juxuny/log-server"
	"github.com/juxuny/log-server/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port     int
	certFile string
	keyFile  string
	logger   = log.NewLogger("[logd]")

	fileLogger           log_server.Logger
	logDir               string
	flushDurationSeconds int
	cacheSize            int
)

func init() {
	flag.StringVar(&certFile, "cert", "", "cert file")
	flag.StringVar(&keyFile, "key", "", "cert key file")
	flag.StringVar(&logDir, "d", "log", "directory for log data")
	flag.IntVar(&flushDurationSeconds, "flush", 30, "flush duration")
	flag.IntVar(&cacheSize, "size", 10000, "cache size")
	flag.IntVar(&port, "p", 40000, "listen port")
}

func main() {
	flag.Parse()
	fileLogger = log_server.NewDefaultFileLogger(logDir, cacheSize, flushDurationSeconds)
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		logger.Error("failed to listen:", err)
	}
	opts := make([]grpc.ServerOption, 0)
	if certFile != "" {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			logger.Error("cannot load TLS credentials: ", err)
			os.Exit(-1)
		}
		opts = append(opts, grpc.Creds(tlsCredentials))
	}
	s := grpc.NewServer(opts...)
	fmt.Println("listen", addr)
	log_server.RegisterLogServerServer(s, &server{})

	go func() {
		if err := s.Serve(ln); err != nil {
			logger.Error("failed to serve:", err)
		}
	}()
	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	finished := make(chan bool)
	go func() {
		s.Stop()
		logger.Info("flush log data")
		if err := fileLogger.Flush(); err != nil {
			logger.Error(err)
		}
		finished <- true
	}()
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-finished:
		logger.Info("server shutdown properly")
	case <-ctx.Done():
		logger.Info("timeout of 15 seconds.")
	}
}
