package main

import (
	"zinx/ZinxDemo/ProtoDemo/pb"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main(){
	person:=&pb.Person{
		Name:"Jack",
		Age:18,
	}
	
	data,err:=proto.Marshal(person)
	if err!=nil{
		fmt.Println("marshal err",err)
		return
	}
	newperson:=&pb.Person{}
	err=proto.Unmarshal(data,newperson)
	if err!=nil{
		fmt.Println("unmarshal err",err)
		return
	}
	fmt.Println(person)
	fmt.Println(newperson)
}
