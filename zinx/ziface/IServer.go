package ziface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器
	Server()
	//添加路由
	AddRouter(routerid uint32,router IRouter)
	Getconnmng()IConnManager
	AddOnConnStart(func(connection IConnection))
	AddOnConnStop(func(connection IConnection))
	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)
}
