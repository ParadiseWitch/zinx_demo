package znet

import (
	"errors"
	"fmt"
	"sync"
	"zinx_demo/ziface"
)

type ConnManager struct {
	connections map[uint32]ziface.IConnection //管理的连接信息
	connLock    sync.RWMutex                  //读写连接的读写锁

}

func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]ziface.IConnection),
	}
}

func (cm *ConnManager) Add(conn ziface.IConnection) {
	//保护共享资源Map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	//将conn连接添加到ConnMananger中
	cm.connections[conn.GetConnID()] = conn
	fmt.Println("connection add to ConnManager successfully: conn num = ", cm.Len())

}
func (cm *ConnManager) Remove(conn ziface.IConnection) {
	//保护共享资源Map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	delete(cm.connections, conn.GetConnID())
	fmt.Println("connection Remove ConnID=", conn.GetConnID(), " successfully: conn num = ", cm.Len())
}
func (cm *ConnManager) Get(connID uint32) (ziface.IConnection, error) {
	//保护共享资源Map 加读锁
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	if conn, ok := cm.connections[connID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

func (cm *ConnManager) Len() int {
	return len(cm.connections)
}

func (cm *ConnManager) ClearConn() {
	//保护共享资源Map 加写锁
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	for connID, conn := range cm.connections {
		conn.Stop()
		delete(cm.connections, connID)
	}
	fmt.Println("Clear All Connections successfully: conn num = ", cm.Len())
}
