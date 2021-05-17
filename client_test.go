package log_server

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	c, err := NewClient("127.0.0.1:40000", "")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 1000*1000; i++ {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()
			_, err = c.Add(ctx, &AddReq{
				App:  "api",
				Data: fmt.Sprintf("[%s] %d", time.Now(), i),
			})
			if err != nil {
				t.Fatal(err)
			}
		}()
	}
}
