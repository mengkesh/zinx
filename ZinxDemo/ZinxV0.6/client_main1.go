package main

import (
	"fmt"
	"time"
	"net"
	"zinx/zinx/znet"
	"io"
)

func main() {
	fmt.Println("client start")
	time.Sleep(time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err,", err)
		return
	}
	//for {
	//	//写
	//	_,err:=conn.Write([]byte("hello zinx"))
	//	if err!=nil{
	//		fmt.Println("write conn err,",err)
	//		return
	//	}
	//	//读
	//	buf:=make([]byte,512)
	//	n,err:=conn.Read(buf)
	//	if err!=nil{
	//		fmt.Println("read buf error")
	//		return
	//	}
	//	fmt.Printf("server call back:%scnt=%d\n",buf,n)
	//	time.Sleep(1 *time.Second)
	//}
	dp := znet.NewDataPack()
	go func() {
		for {
			headbuf := make([]byte, dp.GetHeadLen())
			_, err = io.ReadFull(conn, headbuf)
			if err != nil {
				fmt.Println("read headdata err,", err)
				break
			}
			msg, err := dp.UnPack(headbuf)
			if err != nil {
				fmt.Println("unpack err,", err)
				break
			}
			if msg.GetDataLen() > 0 {
				msg.(*znet.Message).Data = make([]byte, msg.GetDataLen())
				_, err := io.ReadFull(conn, msg.(*znet.Message).Data)
				if err != nil {
					fmt.Println("read headdata err,", err)
					break
				}
			}
			fmt.Println("---> Recv Server Msg : id = ", msg.GetDataID(), "len = ", msg.GetDataLen(), " data = ", string(msg.GetData()))
		}

	}()
	for {

		msg := znet.NewMessage([]byte("hello,zinx"), 1)
		binarydata, err := dp.Pack(msg)
		if err != nil {
			fmt.Println("pack err,", err)
			break
		}
		_, err = conn.Write(binarydata)
		if err != nil {
			fmt.Println("write to server err,", err)
			break
		}

		time.Sleep(time.Second)
	}

}
