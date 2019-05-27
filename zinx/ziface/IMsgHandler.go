package ziface

type IMsgHandler interface {
	AddRouter(msgid uint32,router IRouter)
	DoMsgHandler(request IRequest)
	StartWorkPool()
	SendMsgToTaskQueue(request IRequest)
}