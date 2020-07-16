package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	proto "my/greet/proto"
)

func main() {
	service := micro.NewService(micro.Name("hello.client")) // 客户端服务名称
	service.Init()

	helloservice := proto.NewHelloService("helloSrv", service.Client())
	res, err := helloservice.Ping(context.TODO(), &proto.Request{Name: "World ^_^"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res.Msg)
}
