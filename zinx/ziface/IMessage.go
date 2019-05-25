package ziface

type IMessage interface {
	GetData()[]byte
	GetDataID()uint32
	GetDataLen()uint32
	SetData([]byte)
	SetDataID(uint32)
	SetDataLen(uint32)
}

