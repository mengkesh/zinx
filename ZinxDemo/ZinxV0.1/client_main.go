package main

import (
	"fmt"
	"time"
	"net"
)

func main(){
	fmt.Println("client start")
	time.Sleep(time.Second)
	conn,err:=net.Dial("tcp","127.0.0.1:7777")
	if err!=nil{
		fmt.Println("client start err,",err)
		return
	}
	for {
		//写
		_,err:=conn.Write([]byte("hello zinx"))
		if err!=nil{
			fmt.Println("write conn err,",err)
			return
		}
		//读
		buf:=make([]byte,512)
		n,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back:%scnt=%d\n",buf,n)
		time.Sleep(1 *time.Second)
	}
}
