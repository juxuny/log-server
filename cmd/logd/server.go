package main

import (
	"context"
	"github.com/juxuny/log-server"
)

type server struct {
	log_server.UnimplementedLogServerServer
}

func (t *server) Add(ctx context.Context, req *log_server.AddReq) (resp *log_server.AddResp, err error) {
	resp = &log_server.AddResp{}
	return resp, nil
}

func (t *server) Ping(ctx context.Context, req *log_server.PingReq) (resp *log_server.PingResp, err error) {
	resp = &log_server.PingResp{}
	return resp, nil
}

func (t *server) Info(ctx context.Context, req *log_server.InfoReq) (resp *log_server.InfoResp, err error) {
	resp = &log_server.InfoResp{}
	return resp, nil
}
