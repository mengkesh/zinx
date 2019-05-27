package main

import (
	"zinx/zinx/znet"
	"zinx/zinx/ziface"
	"fmt"
)
type PingRouter struct {
	znet.BaseRouter
}

func (this *PingRouter)Handle(request ziface.IRequest){
	//fmt.Println("Call Router Handle...")
	err:=request.GetConnection().Send([]byte("ping...ping...ping"),1)
	if err!=nil{
		fmt.Println("send err,",err)
		return
		//panic("send data err")
	}
	//_,err:=request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	//if err!=nil{
	//	fmt.Println("call  ping error")
	//}
}
type HelloRouter struct {
	znet.BaseRouter
}

func (this *HelloRouter)Handle(request ziface.IRequest){
	fmt.Println("Call Router Handle...")
	err:=request.GetConnection().Send([]byte("hello...hello...hello"),1)
	if err!=nil{
		fmt.Println("send err,",err)
		return
	}

}

func main(){
	s:=znet.NewServer("zinxV0.6")
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&HelloRouter{})
	s.Server()
	return
}
