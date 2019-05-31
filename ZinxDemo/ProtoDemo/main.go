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
		Emails:[]string{"123@111","234@123"},
		Phone:[]*pb.Phone{
			&pb.Phone{
			Number:"1234567",
			},
			&pb.Phone{
				Number:"34578",
			},
		},
		Data:&pb.Person_School{
			School:"chaunzhiboke",
		},
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
	fmt.Println(newperson.GetName(),newperson.GetAge(),newperson.GetEmails(),newperson.GetPhone(),newperson.GetData())
}
