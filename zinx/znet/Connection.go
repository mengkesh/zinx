package znet

import (
	"net"
	"zinx/ziface"
	"fmt"
	"io"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	//handleAPI ziface.HandleFunc
	Router ziface.IRouter
}
func NewConnection(conn *net.TCPConn,connID uint32,router ziface.IRouter)ziface.IConnection{
	c:=&Connection{
		Conn:conn,
		ConnID:connID,
		isClosed:false,
		//handleAPI:callback_api,
		Router:router,
	}
	return c
}
func(this *Connection)StartReader(){
	fmt.Println("Reader go is startin....")
	defer fmt.Println("connID = ", this.ConnID, "Reader is exit, remote addr is = ", this.GetRemoteAddr().String())
	defer this.Stop()
	for {
		//buf:=make([]byte,config.GlobalObject.MaxPackageSize)
		//cnt,err:=this.Conn.Read(buf)
		//if cnt==0{
		//	fmt.Println("client outline")
		//	break
		//}
		//if err!=nil&&err!=io.EOF{
		//	fmt.Println("recv buf err",err)
		//	continue
		//}
		dp:=NewDataPack()
		datahed:=make([]byte,dp.GetHeadLen())
		n,err:=io.ReadFull(this.Conn,datahed)
		if n==0{
			fmt.Println("client outline")
			break
		}
		if err!=nil&&err!=io.EOF{
			fmt.Println("read datahead err,",err)
			continue
		}
		msg,err:=dp.UnPack(datahed)
		if err!=nil{
			fmt.Println("unpack err,",err)
			break
		}
		if msg.GetDataLen()>0 {
			msg.(*Message).Data=make([]byte,msg.GetDataLen())
			n,err:=io.ReadFull(this.Conn,msg.(*Message).Data)
			if n==0{
				fmt.Println("client outline")
				break
			}
			if err!=nil&&err!=io.EOF{
				fmt.Println("read data err,",err)
				continue
			}
		}
		req:=NewRequest(this,msg)
		go func() {
			this.Router.PreHandle(req)
			//time.Sleep()
			this.Router.Handle(req)
			this.Router.PostHandle(req)
		}()

		/*if err:=this.handleAPI(req);err!=nil{
			fmt.Println("ConnID", this.ConnID, "Handle is error", err)
			break
		}*/
	}
}
func(this *Connection)Start(){
	fmt.Println("Conn Start（）  ... id = ", this.ConnID)
	go this.StartReader()
}
func(this *Connection)Stop(){
	fmt.Println("c. Stop() ... ConnId = ", this.ConnID)
	if this.isClosed==true{
		return
	}
	this.isClosed=true
	_=this.Conn.Close()
}
func(this *Connection)GetConnId()uint32{
return  this.ConnID
}
func(this *Connection)GetTCPConnection()*net.TCPConn{
return this.Conn
}
func(this *Connection)GetRemoteAddr()net.Addr{
return this.Conn.RemoteAddr()
}
func(this *Connection)Send(data []byte,dataid uint32)error{
	dp:=NewDataPack()
	msg:=NewMessage(data,dataid)
	binarydata,err:=dp.Pack(msg)
	if err!=nil{
		fmt.Println("pack err,",err)
		return err
	}
if _,err:=this.Conn.Write(binarydata);err!=nil{
	fmt.Println("send buf error")
	return err
}
return nil
}