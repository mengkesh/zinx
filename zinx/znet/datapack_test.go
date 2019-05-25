package znet

import (
	"testing"
	"net"
	"fmt"
	"io"
)

func TestDataPack(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server create listener err,", err)
		return
	}
	defer listener.Close()
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("listen err,", err)
				break
			}
			go func() {
				dp := NewDataPack()
				buf := make([]byte, dp.GetHeadLen())
				for {
					_, err = io.ReadFull(conn, buf)
					if err != nil {
						fmt.Println("read err,", err)
						break
					}
					msg, err := dp.UnPack(buf)
					if err != nil {
						fmt.Println(err)
						break
					}
					if msg.GetDataLen() > 0 {
						databuf := make([]byte, msg.GetDataLen())
						_, err = io.ReadFull(conn, databuf)
						if err != nil {
							fmt.Println("read err,", err)
							break
						}
						msg.SetData(databuf)
						fmt.Println("------->>data:", string(msg.GetData()), "---->>len:", msg.GetDataLen(), "------->>id", msg.GetDataID())
					}

				}

			}()
		}

	}()
	clintconn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err,", err)
		return
	}
	clientdp := NewDataPack()
	msg1 := &Message{
		Data:    []byte("hello"),
		DataLen: 5,
		DataId:  1,
	}
	msg2 := &Message{
		Data:    []byte("zinx"),
		DataLen: 4,
		DataId:  2,
	}
	data1, err := clientdp.Pack(msg1)
	if err != nil {
		fmt.Println("pack err,", err)
		return
	}
	data2, err := clientdp.Pack(msg2)
	if err != nil {
		fmt.Println("pack err,", err)
		return
	}
	data1=append(data1,data2...)
	clintconn.Write(data1)
	select {

	}
}
