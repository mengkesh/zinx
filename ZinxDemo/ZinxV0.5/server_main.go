package main

import (
	"zinx/znet"
	"zinx/ziface"
	"fmt"
)
type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("Call Router Handle...")
	err:=request.GetConnection().Send([]byte("ping...ping...ping"),1)
	if err!=nil{
		fmt.Println("send err,",err)
	}
	//_,err:=request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	//if err!=nil{
	//	fmt.Println("call  ping error")
	//}
}

func main(){
	s:=znet.NewServer("zinx v0.1")
	s.AddRouter(&PingRouter{})
	s.Server()
	return
}
