package apis

import (
	"zinx/zinx/znet"
	"zinx/zinx/ziface"
	"zinx/mmo_game_server/pb"
	"github.com/golang/protobuf/proto"
	"fmt"
	"zinx/mmo_game_server/core"
)

type WorldChat struct {
	znet.BaseRouter
}
func(wc *WorldChat)Handle(request ziface.IRequest){
	proto_msg:=&pb.Talk{}
	err:=proto.Unmarshal(request.GetMsg().GetData(),proto_msg)
	if err!=nil{
		fmt.Println("Talk message unmarshal err,",err)
		return
	}
	pid,err:=request.GetConnection().GetProperty("pid")
	if err!=nil{
		fmt.Println("get Pid err,",err)
		return
	}
	player:=core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.SendTalkMsgToAll(proto_msg.Content)
}