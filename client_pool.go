package log_server

import (
	"context"
	"fmt"
	"github.com/juxuny/log-server/log"
	"github.com/pkg/errors"
	"math/rand"
	"os"
	"runtime/debug"
	"sync"
	"time"
)

var logger = log.NewLogger("[LOG]")

type clientPool struct {
	list    []*client
	rand    *rand.Rand
	reqPool *sync.Pool
}

func NewClientPool(cert string, host ...string) (*clientPool, error) {
	pool := &clientPool{}
	pool.rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	pool.list = make([]*client, len(host))
	var err error
	for i := 0; i < len(host); i++ {
		pool.list[i], err = NewClient(host[i], cert)
		if err != nil {
			return nil, errors.Wrap(err, "init client pool failed")
		}
	}
	pool.reqPool = &sync.Pool{
		New: func() interface{} {
			return &AddReq{}
		},
	}
	return pool, nil
}

func (t *clientPool) Add(ctx context.Context, app string, data string) error {
	req := t.reqPool.Get().(*AddReq)
	req.App = app
	req.Data = data
	defer func() {
		t.reqPool.Put(req)
	}()
	defer func() {
		if err := recover(); err != nil {
			_, _ = os.Stderr.Write([]byte(fmt.Sprint(err)))
			debug.PrintStack()
		}
	}()
	var err error
	for i := 0; i < 3; i++ {
		r := t.rand.Intn(len(t.list))
		_, err = t.list[r].Add(ctx, req)
		if err == nil {
			break
		}
	}
	return err
}
