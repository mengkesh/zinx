package znet

import "zinx/zinx/ziface"

type Message struct {
	Data []byte
	DataId uint32
	DataLen uint32
}
func NewMessage(data []byte,dataid uint32)ziface.IMessage{
	m:=&Message{
		Data:data,
		DataId:dataid,
		DataLen:uint32(len(data)),
	}
	return m
}
func(this *Message)GetData()[]byte{
	return this.Data
}
func(this *Message)GetDataID()uint32{
	return this.DataId
}
func(this *Message)GetDataLen()uint32{
	return this.DataLen
}
func(this *Message)SetData(data []byte){
	this.Data=data
}
func(this *Message)SetDataID(dataid uint32){
	this.DataId=dataid
}
func(this *Message)SetDataLen(datalen uint32){
	this.DataLen=datalen
}
