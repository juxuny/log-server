package log_server

import (
	"context"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("127.0.0.1:40000", "cert/ca-cert.pem")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	_, err = c.Add(ctx, &AddReq{
		App:  "api",
		Data: "100",
	})
	if err != nil {
		t.Fatal(err)
	}
}
