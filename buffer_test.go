package log_server

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNewLogBuffer(t *testing.T) {
	tmp := bytes.NewBuffer(nil)
	buf := NewLogBuffer(100000, time.Second*3, func(data []byte) error {
		tmp.Write(data)
		if _, err := os.Stdout.Write(data); err != nil {
			t.Fatal(err)
		}
		fmt.Println(time.Now())
		return nil
	})
	wg := sync.WaitGroup{}
	groupNum := 100
	startTime := time.Now()
	for i := 0; i < groupNum; i++ {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				time.Sleep(time.Second * 4)
				_ = buf.Append(fmt.Sprintf("%d\n", start+j))
			}
		}(i * 10)
	}
	wg.Wait()
	_ = buf.Flush()
	l := strings.Split(strings.Trim(tmp.String(), "\n"), "\n")
	t.Log(len(l))
	if len(l) != groupNum*10 {
		t.Fatal("wrong line number", len(l))
	}
	t.Log("waste time: ", time.Now().Sub(startTime))
}
