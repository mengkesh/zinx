package apis

import (
	"zinx/zinx/znet"
	"zinx/zinx/ziface"
	"zinx/mmo_game_server/pb"
	"github.com/golang/protobuf/proto"
	"zinx/mmo_game_server/core"
	"fmt"
)

type Move struct {
	znet.BaseRouter
}
func(m *Move)Handle(request ziface.IRequest){
	proto_msg:=&pb.Position{}
	proto.Unmarshal(request.GetMsg().GetData(),proto_msg)
	pid,_:=request.GetConnection().GetProperty("pid")
	fmt.Println("player id = ",pid.(int32), " move --> ", proto_msg.X, ", ", proto_msg.Z, ", ", proto_msg.V)
	player:=core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.ShowHide(proto_msg.X,proto_msg.Y,proto_msg.Z,proto_msg.V)
	player.UpdatePosdtion(proto_msg.X,proto_msg.Y,proto_msg.Z,proto_msg.V)

}