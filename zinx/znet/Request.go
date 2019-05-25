package znet

import "zinx/ziface"

type Request struct {
	conn ziface.IConnection
	//data []byte
	//len int
	msg ziface.IMessage
}
func NewRequest(conn ziface.IConnection,msg ziface.IMessage)ziface.IRequest{
	req:=&Request{
		conn:conn,
		//data:data,
		//len:len,
		msg:msg,
	}
	return req
}
func(this *Request)GetConnection()ziface.IConnection{
	return this.conn
}
//func(this *Request)GetData()[]byte{
//	return this.data
//}
//func(this *Request)GetDataLen()int{
//	return this.len
//}
func(this *Request)GetMsg()ziface.IMessage{
	return this.msg
}