package ziface

type IMessage interface {
	//获取消息的Id
	GetMsgId() uint32
	//获取消息的长度
	GetMsgDataLen() uint32
	//获取消息的内容
	GetMsgData() []byte
	//设置消息的Id
	SetMsgId(uint32)
	//设置消息的长度
	SetMsgDataLen(uint32)
	//设置消息的内容
	SetMsgData([]byte)
}
