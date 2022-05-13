package znet

import (
	"zinx_demo/ziface"
)

type Request struct {
	conn ziface.IConnection
	msg  ziface.IMessage //客户端请求的数据
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

//获取请求的消息的ID
func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
