package znet

import (
	"net"
	"zinx/zinx/ziface"
	"fmt"
	"io"
	"errors"
	"strconv"
	"zinx/zinx/config"
)

type Connection struct {
	server   ziface.IServer
	Conn     *net.TCPConn
	ConnID   uint32
	isClosed bool
	//handleAPI ziface.HandleFunc
	MsgHandler ziface.IMsgHandler
	msgchan    chan []byte
	isquit     chan bool
}

func NewConnection(server ziface.IServer, conn *net.TCPConn, connID uint32, handler ziface.IMsgHandler) ziface.IConnection {
	c := &Connection{
		server:   server,
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		//handleAPI:callback_api,
		MsgHandler: handler,
		msgchan:    make(chan []byte),
		isquit:     make(chan bool),
	}
	c.server.Getconnmng().Add(c)
	return c
}
func (this *Connection) StartReader() {
	fmt.Println("Reader go is start....")
	defer fmt.Println("start reader is stop,connID = ", this.ConnID, "Reader is exit, remote addr is = ", this.GetRemoteAddr().String())
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
		dp := NewDataPack()
		datahed := make([]byte, dp.GetHeadLen())
		n, err := io.ReadFull(this.Conn, datahed)
		if n <= 0 {
			fmt.Println("client outline")
			return
		}
		if err != nil && err != io.EOF {
			fmt.Println("read datahead err,", err)
			return
		}
		msg, err := dp.UnPack(datahed)
		if err != nil {
			fmt.Println("unpack err,", err)
			return
		}
		if msg.GetDataLen() > 0 {
			msg.(*Message).Data = make([]byte, msg.GetDataLen())
			n, err := io.ReadFull(this.Conn, msg.(*Message).Data)
			if n == 0 {
				fmt.Println("client outline")
				break
			}
			if err != nil && err != io.EOF {
				fmt.Println("read data err,", err)
				break
			}
		}
		req := NewRequest(this, msg)
		if config.GlobalObject.WorkPoolSize > 0 {
			this.MsgHandler.SendMsgToTaskQueue(req)
		} else {
			go this.MsgHandler.DoMsgHandler(req)
		}

		/*if err:=this.handleAPI(req);err!=nil{
			fmt.Println("ConnID", this.ConnID, "Handle is error", err)
			break
		}*/
	}
}
func (this *Connection) StartWriter() {
	fmt.Println("writer is start")
	defer fmt.Println("start writer is stop")
	for {
		select {
		case data := <-this.msgchan:
			_, err := this.Conn.Write(data)
			if err != nil {
				fmt.Println("Send err in startwriter,", err)
				return
			}
		case <-this.isquit:
			return
		}
	}

}
func (this *Connection) Start() {
	fmt.Println("Conn Start（）  ... id = ", this.ConnID)
	go this.StartReader()
	go this.StartWriter()
	this.server.CallOnConnStart(this)
}
func (this *Connection) Stop() {
	fmt.Println("c. Stop() ... ConnId = ", this.ConnID)
	this.server.CallOnConnStop(this)
	if this.isClosed == true {
		return
	}
	this.isClosed = true
	this.isquit <- true
	_ = this.Conn.Close()
	this.server.Getconnmng().Remove(this.ConnID)
	close(this.msgchan)
	close(this.isquit)
}
func (this *Connection) GetConnId() uint32 {
	return this.ConnID
}
func (this *Connection) GetTCPConnection() *net.TCPConn {
	return this.Conn
}
func (this *Connection) GetRemoteAddr() net.Addr {
	return this.Conn.RemoteAddr()
}
func (this *Connection) Send(data []byte, dataid uint32) error {
	if this.isClosed == true {
		return errors.New("connection is closed, conID is :" + strconv.Itoa(int(this.GetConnId())))
	}

	dp := NewDataPack()
	msg := NewMessage(data, dataid)
	binarydata, err := dp.Pack(msg)
	if err != nil {
		fmt.Println("pack err,", err)
		return err
	}
	//if _, err := this.Conn.Write(binarydata); err != nil {
	//	fmt.Println("send buf err", err)
	//	return err
	//}
	this.msgchan <- binarydata
	return nil
}
