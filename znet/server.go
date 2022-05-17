package znet

import (
	"fmt"
	"net"
	"time"
	"zinx_demo/utils"
	"zinx_demo/ziface"
)

type Server struct {
	Name        string                        //服务器的名称
	IPVersion   string                        //tcp4 or other
	IP          string                        //服务绑定的IP地址
	Port        int                           //服务绑定的端口
	msgHandler  ziface.IMsgHandler            //当前Server的消息管理模块，用来绑定MsgId和对应的处理方法
	ConnMgr     ziface.IConnManager           //当前Server的链接管理器
	OnConnStart func(conn ziface.IConnection) //该Server的连接创建时Hook函数
	OnConnStop  func(conn ziface.IConnection) //该Server的连接断开时的Hook函数
}

func NewServer(name string) ziface.IServer {
	utils.GlobalObject.Reload()
	return &Server{
		Name:       name,
		IPVersion:  "tcp4",
		IP:         "0.0.0.0",
		Port:       7777,
		msgHandler: NewMsgHandler(),
		ConnMgr:    NewConnManager(), //创建ConnManager
	}
}

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
}

func (s *Server) Start() {
	fmt.Printf("[START] Server listenner at IP: %s, Port %d, is starting\\n\n", s.IP, s.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)
	go func() {
		//0 启动worker工作池机制
		s.msgHandler.StartWorkerPool()
		//1 获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}
		fmt.Println("start Zinx server  ", s.Name, " succ, now listenning...")
		var cid uint32 = 0
		for {
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			// 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			dealConn := NewConntion(s, conn, cid, s.msgHandler)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)

	//将其他需要清理的连接信息或者其他信息 也要一并停止或者清理
	s.ConnMgr.ClearConn()
}

func (s *Server) Serve() {
	s.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

//设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

//设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

//调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

//调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}
