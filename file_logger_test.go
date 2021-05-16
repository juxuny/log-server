package log_server

import (
	"fmt"
	"path"
	"testing"
	"time"
)

type TestGenerator struct {
	dir string
}

func (t *TestGenerator) Gen(app string) string {
	return path.Join(t.dir, app+"_"+time.Now().Format("20060102_1504")+".log")
}

func TestNewFileLogger(t *testing.T) {
	logger := NewFileLogger("tmp", 10, 10, &TestGenerator{dir: "tmp"})
	for i := 0; i < 1000*1000; i++ {
		if err := logger.Info("api", fmt.Sprintf("%d\n", i)); err != nil {
			fmt.Printf("%+v", err)
			t.Fatal()
		}
		time.Sleep(time.Second)
	}
}
