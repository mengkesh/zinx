package znet

import (
	"zinx/zinx/ziface"
	"bytes"
	"encoding/binary"
)

type DataPack struct {
}
func NewDataPack()ziface.IDataPack{
	return &DataPack{}
}
func (this *DataPack)GetHeadLen()uint32{
	return 8
}
func (this *DataPack)Pack(message ziface.IMessage)([]byte,error){
	databuffer:=bytes.NewBuffer([]byte{})
	err:=binary.Write(databuffer,binary.LittleEndian,message.GetDataLen())
	if err!=nil{
		return nil,err
	}
	err=binary.Write(databuffer,binary.LittleEndian,message.GetDataID())
	if err!=nil{
		return nil,err
	}
	err=binary.Write(databuffer,binary.LittleEndian,message.GetData())
	if err!=nil{
		return nil,err
	}
	return databuffer.Bytes(),nil
}
func(this *DataPack)UnPack(data []byte)(ziface.IMessage,error){
	message:=&Message{}
	databuffer:=bytes.NewBuffer(data)
	err:=binary.Read(databuffer,binary.LittleEndian,&message.DataLen)
	if err!=nil{
		return nil,err
	}
	err=binary.Read(databuffer,binary.LittleEndian,&message.DataId)
	if err!=nil{
		return nil,err
	}
	return message,nil
}