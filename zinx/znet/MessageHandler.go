package znet

import (
	"zinx/zinx/ziface"
	"fmt"
	"zinx/zinx/config"
)

type MessageHandler struct {
	Apis map[uint32]ziface.IRouter
	TaskQueue []chan ziface.IRequest
	WorkPool uint32
}
func NewMessageHandler()ziface.IMsgHandler{
	return &MessageHandler{
		Apis:make(map[uint32]ziface.IRouter),
		TaskQueue:make([]chan ziface.IRequest,config.GlobalObject.WorkPoolSize),
		WorkPool:config.GlobalObject.WorkPoolSize,
	}
}
func (mh *MessageHandler)AddRouter(msgid uint32,router ziface.IRouter){
	if _,ok:=mh.Apis[msgid];ok{
		fmt.Println("repeat Api msgid=",msgid)
		return
	}
	mh.Apis[msgid]=router
	fmt.Println("Add APi msgid=",msgid,"success")
}
func(this *MessageHandler)DoMsgHandler(request ziface.IRequest){
	router,ok:=this.Apis[request.GetMsg().GetDataID()]
	if !ok{
		fmt.Println("Api msgid=",request.GetMsg().GetDataID(),"is not found!need add")
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}
func(this *MessageHandler)StartOneWork(i int,taskqueue chan ziface.IRequest){
	fmt.Println("worker-",i,"-is start")
	for  {
		select {
		case request := <-taskqueue:
			this.DoMsgHandler(request)
		}
	}
}
func(this *MessageHandler)StartWorkPool(){
	for i:=0;i<int(this.WorkPool);i++{
		this.TaskQueue[i]=make(chan ziface.IRequest,config.GlobalObject.MaxWorkerTaskLen)
		go this.StartOneWork(i,this.TaskQueue[i])
	}
}
func(this *MessageHandler)SendMsgToTaskQueue(request ziface.IRequest){
	workid:=request.GetConnection().GetConnId()%this.WorkPool
	this.TaskQueue[workid]<-request
}