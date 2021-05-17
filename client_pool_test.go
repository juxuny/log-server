package log_server

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestNewClientPool(t *testing.T) {
	cp, err := NewClientPool("", "127.0.0.1:40000")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10000*1000; i++ {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
			defer cancel()
			if err := cp.Add(ctx, "api", fmt.Sprintf("[%s] %d", time.Now().Format("2006-01-02 15:04:05"), i)); err != nil {
				t.Fatal(err)
			}
		}()
	}
}
