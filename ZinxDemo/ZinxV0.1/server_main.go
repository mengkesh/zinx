package main

import (
	"zinx/znet"
	"zinx/ziface"
	"fmt"
)
type PingRouter struct {
	znet.BaseRouter
}
func(this *PingRouter)PreHandle(request ziface.IRequest){
	fmt.Println("Call Router PreHandle...")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
		if err != nil {
			fmt.Println("call back before ping error")
		}

}
func (this *PingRouter)Handle(request ziface.IRequest){
	fmt.Println("Call Router Handle...")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	if err!=nil{
		fmt.Println("call  ping error")
	}
}
func(this *PingRouter)PostHandle(request ziface.IRequest){
	fmt.Println("Call Router PostHandle")
	_,err:=request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err!=nil{
		fmt.Println("call back after ping error")
	}
}
func main(){
	s:=znet.NewServer("zinx v0.1")
	s.AddRouter(&PingRouter{})
	s.Server()
	return
}
