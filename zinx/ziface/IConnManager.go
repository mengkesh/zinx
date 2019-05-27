package ziface

type IConnManager interface {
	Add(conn IConnection)
	Remove(connId uint32)
	Get(ConnID uint32)(IConnection,error)
	Len()uint32
	Clearconn()
}