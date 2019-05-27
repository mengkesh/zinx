package znet

import (
	"zinx/zinx/ziface"
	"sync"
	"fmt"
	"errors"
)

type ConnManager struct {
	connectons map[uint32]ziface.IConnection
	connLock   sync.RWMutex
}

func NewConneManager() ziface.IConnManager {
	return &ConnManager{
		connectons: make(map[uint32]ziface.IConnection),
	}
}
func (this *ConnManager) Add(conn ziface.IConnection) {
	this.connLock.Lock()
	defer this.connLock.Unlock()
	this.connectons[conn.GetConnId()] = conn
	fmt.Println("Add Conn success,id=", conn.GetConnId())
}
func (this *ConnManager) Remove(connId uint32) {
	this.connLock.Lock()
	defer this.connLock.Unlock()
	delete(this.connectons, connId)
	fmt.Println("Remove connid = ", connId, " from manager succ!!")
}
func (this *ConnManager) Get(ConnID uint32) (ziface.IConnection, error) {
	this.connLock.RLock()
	defer this.connLock.RUnlock()

	if conn, ok := this.connectons[ConnID]; ok {
		return conn, nil
	} else {
		return nil, errors.New("conn is not found")
	}
}
func (this *ConnManager) Len() uint32 {
	return uint32(len(this.connectons))
}
func (this *ConnManager) Clearconn() {
	this.connLock.Lock()
	defer this.connLock.Unlock()
	for connid,conn:=range this.connectons{
		conn.Stop()
		delete(this.connectons,connid)
	}
}
