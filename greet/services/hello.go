package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	proto "my/greet/proto"
)

type Hello struct{}

func (h *Hello) Ping(ctx context.Context, req *proto.Request, res *proto.Response) error {
	res.Msg = "Hello " + req.Name
	return nil
}

func main() {
	registry := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("helloSRV"), // 服务名称
		micro.Registry(registry),
	)
	service.Init()
	if err := proto.RegisterHelloHandler(service.Server(), new(Hello)); err != nil {
		panic(err)
	}
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
