package log_server

import (
	"context"
	"github.com/pkg/errors"
	"math/rand"
	"sync"
	"time"
)

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
	r := t.rand.Intn(len(t.list))
	req := t.reqPool.Get().(*AddReq)
	req.App = app
	req.Data = data
	defer func() {
		t.reqPool.Put(req)
	}()
	_, err := t.list[r].Add(ctx, req)
	return err
}
