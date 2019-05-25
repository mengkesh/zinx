package znet

import (
	"zinx/zinx/ziface"
	"fmt"
)

type MessageHandler struct {
	Apis map[uint32]ziface.IRouter
}
func NewMessageHandler()ziface.IMsgHandler{
	return &MessageHandler{
		Apis:make(map[uint32]ziface.IRouter),
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