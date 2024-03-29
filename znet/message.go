package znet

type Message struct {
	Id      uint32 //消息的ID
	DataLen uint32 //消息的长度
	Data    []byte //消息的内容
}

func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetMsgData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetMsgDataLen(len uint32) {
	m.DataLen = len
}

func (m *Message) SetMsgData(data []byte) {
	m.Data = data
}
