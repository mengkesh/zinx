package main

import (
	"zinx/zinx/znet"
	"zinx/zinx/ziface"
	"fmt"
	"zinx/mmo_game_server/core"
	"zinx/mmo_game_server/apis"
)

func OnConnectionAdd(conn ziface.IConnection){
	fmt.Println("conn add......")
	player:=core.NewPlayer(conn)
	player.ReturnPid()
	player.ReturnPlayerPosition()
	core.WorldMgrObj.AddPlayer(player)
	player.SyncSurrounding()
	conn.SetProperty("pid",player.Pid)
	fmt.Println("----> player ID = ", player.Pid, "Online...", ", Player num = ", len(core.WorldMgrObj.Players))
}
func OnConnectionStop(conn ziface.IConnection){
	pid,err:=conn.GetProperty("pid")
	if err!=nil{
		fmt.Println("get pid err in onconnectionstop",err)
		return
	}
	player:=core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.OffLine()
}
func main(){
	s:=znet.NewServer("MMO Game Server")
	s.AddOnConnStart(OnConnectionAdd)
	s.AddOnConnStop(OnConnectionStop)
	s.AddRouter(2,&apis.WorldChat{})
	s.AddRouter(3,&apis.Move{})
	s.Server()
}
