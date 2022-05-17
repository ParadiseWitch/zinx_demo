package znet

import (
	"fmt"
	"net"
	"time"
	"zinx_demo/utils"
	"zinx_demo/ziface"
)

type Server struct {
	Name       string
	IPVersion  string
	IP         string
	Port       int
	msgHandler ziface.IMsgHandler
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
			dealConn := NewConntion(conn, cid, s.msgHandler)
			cid++
			go dealConn.Start()
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

func (s *Server) Serve() {
	s.Start()
	for {
		time.Sleep(10 * time.Second)
	}
}

func NewServer(name string) ziface.IServer {
	utils.GlobalObject.Reload()
	return &Server{
		Name:       name,
		IPVersion:  "tcp4",
		IP:         "0.0.0.0",
		Port:       7777,
		msgHandler: NewMsgHandler(),
	}
}
