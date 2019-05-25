package znet

import (
	"zinx/ziface"
	"fmt"
	"net"
	"zinx/config"
)

type Server struct {
	//服务器ip
	IPVersion string
	IP        string
	//服务器port
	Port int
	//服务器名称
	Name string
	Router ziface.IRouter
	//ziface.IServer
}

/*//定义一个 具体的回显业务 针对type HandleFunc func(*net.TCPConn,[]byte,int) error
func CallBackBusi(request ziface.IRequest) error {
	//回显业务
	fmt.Println("【conn Handle】 CallBack..")
	c := request.GetConnection().GetTCPConnection()
	buf := request.GetData()
	cnt := request.GetDataLen()
	if _, err := c.Write(buf[:cnt]);err !=nil {
		fmt.Println("write back err ", err)
		return err
	}

	return nil
}*/
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      config.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        config.GlobalObject.Host,
		Port:      config.GlobalObject.Port,
		Router:nil,
	}
	return s
}
//启动服务器
//原生socket 服务器编程
func (s *Server) Start() {
	fmt.Printf("[start]Server Listener at IP:%s,Port:%d,is starting..\n", s.IP, s.Port)
	//1 创建套接字  ：得到一个TCP的addr
	addr,err:=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err!=nil{
		fmt.Println("resolve tcp addr error:",err)
		return
	}
	//2 监听服务器地址
	listener,err:=net.ListenTCP(s.IPVersion,addr)
	if err!=nil{
		fmt.Println("listen",s.IPVersion,"err,",err)
		return
	}
	//生成id累加器
	var cid uint32
	cid=0
	//3 阻塞等待客户端发送请求，
	go func() {
		for{
			//阻塞等待客户端请求,
			conn,err:=listener.AcceptTCP()
			if err!=nil{
				fmt.Println("Accept err,",err)
				continue
			}
			//此时conn就已经和对端客户端连接
			go func() {
				//4 客户端有数据请求，处理客户端业务(读、写)
				dealconn:=NewConnection(conn,cid,s.Router)
				cid++
				go dealconn.Start()
			}()

		}
	}()


}
//停止服务器
func (s *Server) Stop() {
	//TODO 将一些服务器资源进行回收...
}
func (s *Server) Server() {
	//启动server的监听功能
	s.Start()//并不希望他永久的阻塞

	//TODO  做一些其他的扩展
	//阻塞//告诉CPU不再需要处理的，节省cpu资源
	select {

	}
}
func(s *Server)AddRouter(router ziface.IRouter){
	s.Router=router
}
