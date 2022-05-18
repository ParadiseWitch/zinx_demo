package ziface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
	SendMsg(msgId uint32, data []byte) error     //直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendBuffMsg(msgId uint32, data []byte) error //直接将Message数据发送给远程的TCP客户端(有缓冲)
	SetProperty(key string, val interface{})     //设置链接属性
	GetProperty(key string) (interface{}, error) //获取链接属性
	RemoveProperty(key string)                   //删除链接属性
}

type HandFunc func(*net.TCPConn, []byte, int) error
