package znet

import (
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	handleAPI ziface.HandleFunc
}
func NewConnection(conn *net.TCPConn,connID uint32,callback_api ziface.HandleFunc)ziface.IConnection{
	c:=&Connection{
		Conn:conn,
		ConnID:connID,
		isClosed:false,
		handleAPI:callback_api,
	}
	return c
}
func(this *Connection)Start(){

}
func(this *Connection)Stop(){

}
func(this *Connection)GetConnId()uint32{
return  0
}
func(this *Connection)GetTCPConnection()*net.TCPConn{
return nil
}
func(this *Connection)GetRemoteAddr()net.Addr{
return nil
}
func(this *Connection)Send(data []byte
)error{
return nil
}