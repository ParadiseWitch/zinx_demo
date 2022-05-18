package main

import (
	"fmt"
	"zinx_demo/ziface"
	"zinx_demo/znet"
)

type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

func (hw *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloZinxRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.8"))
	if err != nil {
		fmt.Println(err)
	}
}

//创建连接的时候执行
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnecionBegin is Called ... ")
	//=============设置两个链接属性，在连接创建之后===========
	fmt.Println("Set conn Name, Home done!")
	conn.SetProperty("Name", "Maid")
	conn.SetProperty("Age", "14")
	//===================================================
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

//连接断开的时候执行
func DoConnectionLost(conn ziface.IConnection) {
	//============在连接销毁之前，查询conn的Name，Home属性=====
	if name, err := conn.GetProperty("Name"); err == nil {
		fmt.Println("Conn Property Name = ", name)
	}
	if age, err := conn.GetProperty("Age"); err == nil {
		fmt.Println("Conn Property Age = ", age)
	}
	//===================================================
	fmt.Println("DoConneciotnLost is Called ... ")
}

func main() {
	s := znet.NewServer("[zinx V0.8]")

	//注册链接hook回调函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
