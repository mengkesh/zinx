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
	//fmt.Println("Call Router Handle...")
	err:=request.GetConnection().Send([]byte("hello...hello...hello"),1)
	if err!=nil{
		fmt.Println("send err,",err)
		return
	}

}
func onconnstart(conn ziface.IConnection){
	fmt.Println("onconnstart is begin now")
	err:=conn.Send([]byte("welcome to zinx"),202)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println("---> Set conn property <---")
	conn.SetProperty("Name","Go3")
	conn.SetProperty("Address","TBD")
	conn.SetProperty("Time","2019-12-12")
}
func onconnstop(conn ziface.IConnection){
	fmt.Println("onconnstop is begin now")
	name,err:=conn.GetProperty("Name")
	if err==nil{
		fmt.Println("Name =", name)
	}
	address,err:=conn.GetProperty("Address")
	if err==nil{
		fmt.Println("Address =", address)
	}
	time,err:=conn.GetProperty("Time")
	if err==nil{
		fmt.Println("Time =", time)
	}
}
func main(){
	s:=znet.NewServer("zinxV0.6")
	s.AddOnConnStart(onconnstart)
	s.AddOnConnStop(onconnstop)
	s.AddRouter(0,&PingRouter{})
	s.AddRouter(1,&HelloRouter{})
	s.Server()
	return
}
