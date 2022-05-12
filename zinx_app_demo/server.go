package main

import (
	"fmt"
	"zinx_demo/ziface"
	"zinx_demo/znet"
)

type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

func (pr *PingRouter) PreHandle(req ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

func main() {
	s := znet.NewServer("[zinx V0.3]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
