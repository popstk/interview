package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/etcdv3"
	proto "my/greet/proto"
)

func main() {
	registry := etcdv3.NewRegistry()
	service := micro.NewService(
		micro.Name("hello.client"),// 客户端服务名称
		micro.Registry(registry))
	service.Init()

	helloService := proto.NewHelloService("helloSRV", service.Client())
	res, err := helloService.Ping(context.TODO(), &proto.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.Msg)
	select {
	}
}
