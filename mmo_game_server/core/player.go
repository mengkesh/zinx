package core

import (
	"zinx/zinx/ziface"
	"sync"
	"math/rand"
	"github.com/golang/protobuf/proto"
	"fmt"
	"zinx/mmo_game_server/pb"
)

type player struct {
	Pid  int32
	Conn ziface.IConnection
	X    float32
	Y    float32
	Z    float32
	V    float32
}

var PidGen int32 = 1
var IdLock sync.Mutex

func NewPlayer(conn ziface.IConnection) *player {
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()
	p := &player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(140 + rand.Intn(10)),
		V:    0,
	}
	return p
}

//玩家可以和对端客户端发送消息的方法
func (p *player) SendMsg(msgid uint32, message proto.Message) error {
	binary_proto_data, err := proto.Marshal(message)
	if err != nil {
		fmt.Println("marshal proto struct err,", err)
		return err
	}
	err = p.Conn.Send(binary_proto_data, msgid)
	if err != nil {
		fmt.Println("player send err,", err)
		return err
	}
	return nil
}
func (p *player) ReturnPid() {
	proto_msg := &pb.SyncPid{
		Pid: p.Pid,
	}
	p.SendMsg(1, proto_msg)
}
func (p *player) ReturnPlayerPosition() {
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(200, proto_msg)
}
func (p *player) SendTalkMsgToAll(content string) {
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	players := WorldMgrObj.GetAllPlayers()
	for _, v := range players {
		v.SendMsg(200, proto_msg)
	}
}
func (p *player) GetSurroundingPlayers() []*player {
	pids := WorldMgrObj.AoiMgr.GetSurroundPIDsByPos(p.X, p.Z)
	players := make([]*player, 0, len(pids))
	for _, v := range pids {
		p := WorldMgrObj.GetPlayerByPid(int32(v))
		players = append(players, p)
	}
	return players
}
func (p *player) SyncSurrounding() {
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	players := p.GetSurroundingPlayers()
	for _, v := range players {
		v.SendMsg(200, proto_msg)
	}
	players_proto_msg := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		//制作一个message Player 消息
		p := &pb.Player{
			Pid:player.Pid,
			P:&pb.Position{
				X:player.X,
				Y:player.Y,
				Z:player.Z,
				V:player.V,
			},
		}

		players_proto_msg = append(players_proto_msg, p)
	}
	//创建一个 Message SyncPlayers
	syncPlayers_proto_msg := &pb.SyncPlayers{
		Ps: players_proto_msg[:],
	}
	//将当前的周边的全部的玩家信息 发送给当前的客户端
	p.SendMsg(202,syncPlayers_proto_msg)
}
func (p *player) UpdatePosdtion(x, y, z, v float32) {
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v
	proto_msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4,
		Data: &pb.BroadCast_P{
			&pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	players := p.GetSurroundingPlayers()
	proto_players := make([]*pb.Player, 0, len(players))
	for _, v := range players {
		v.SendMsg(200, proto_msg)
		proto_players = append(proto_players,
			&pb.Player{
				Pid: v.Pid,
				P: &pb.Position{
					X: v.X,
					Y: v.Y,
					Z: v.Z,
					V: v.V},
			})
	}
	proto_msg1 := &pb.SyncPlayers{
		Ps: proto_players,
	}
	p.SendMsg(202, proto_msg1)
}
func (p *player) OffLine() {
	proto_msg := &pb.SyncPid{
		Pid: p.Pid,
	}
	players := WorldMgrObj.GetAllPlayers()
	for _, v := range players {
		v.SendMsg(201, proto_msg)
	}
	WorldMgrObj.RemovePlayerByPid(p.Pid)
}
func (p *player) ShowHide(x, y, z, v float32) {
	gidbefore := WorldMgrObj.AoiMgr.GetGidByPos(p.X, p.Z)
	/*p.X=x
	p.Y=y
	p.Z=z
	p.V=v*/
	gidafter := WorldMgrObj.AoiMgr.GetGidByPos(x, z)
	if gidbefore != gidafter {
		//将玩家id从老格子删除，在新格子添加
		WorldMgrObj.AoiMgr.RemovePidFromGrid(int(p.Pid),gidbefore)
		WorldMgrObj.AoiMgr.AddPidToGrid(int(p.Pid),gidafter)
		//获取之前格子id
		gridsbefore := WorldMgrObj.AoiMgr.GetSurroundGridsByGid(gidbefore)
		//获取之后格子id
		gridsafter := WorldMgrObj.AoiMgr.GetSurroundGridsByGid(gidafter)
		//存放老格子信息的map
		mgridbefore := make(map[int]*Grid)
		//存放新格子信息的map
		mgridafter := make(map[int]*Grid)
		//组建存放旧格子信息的map
		for _, v := range gridsbefore {
			mgridbefore[v.GID] = v
		}
		var playeridsnew []int
		//组建存放新格子信息的map
		//获取没在旧格子中的新格子玩家id
		for _, v := range gridsafter {
			mgridafter[v.GID] = v
			if _, ok := mgridbefore[v.GID]; !ok {
				//得到新格子中的玩家id
				playeridsnew = append(playeridsnew, v.GetPlayerID()...)
			}
		}
		//封装新格子中的玩家视野中出现移动中玩家的消息
		proto_msg := &pb.BroadCast{
			Pid: p.Pid,
			Tp:  2,
			Data: &pb.BroadCast_P{
				P: &pb.Position{
				X:x,
				Y:y,
				Z:z,
				V:v,
				},
			},
		}
		//遍历新格子玩家id，让移动中玩家出现在新格子中的玩家视野中，新格子中的玩家出现在移动中玩家视野中
		for _,pid:=range playeridsnew{
			WorldMgrObj.Players[int32(pid)].SendMsg(200,proto_msg)
			//封装移动中玩家视野中出现新格子中的玩家的消息
			proto_msgnew:=&pb.BroadCast{
				Pid: WorldMgrObj.Players[int32(pid)].Pid,
				Tp:  2,
				Data: &pb.BroadCast_P{
					P: &pb.Position{
						X:WorldMgrObj.Players[int32(pid)].X,
						Y:WorldMgrObj.Players[int32(pid)].Y,
						Z:WorldMgrObj.Players[int32(pid)].Z,
						V:WorldMgrObj.Players[int32(pid)].V,
					},
				},
			}
			p.SendMsg(200,proto_msgnew)
		}

		var playeridsold []int
		//获取没在新格子中的旧格子玩家id
		for _,v:=range gridsbefore{
			if _,ok:=mgridafter[v.GID];!ok{
				//得到旧格子中的玩家id
				playeridsold=append(playeridsold,v.GetPlayerID()...)
			}
		}
		//封装旧格子中的玩家视野中消失移动中玩家的消息
		proto_msg1:=&pb.SyncPid{
			Pid:p.Pid,
		}
		//遍历旧格子玩家id，让移动中玩家消失在旧格子中的玩家视野中，旧格子中的玩家消失在移动中玩家视野中
		for _,pid:=range playeridsold{
			WorldMgrObj.Players[int32(pid)].SendMsg(201,proto_msg1)
			//封装移动中玩家视野中消失旧格子中的玩家的消息
			proto_msgold:=&pb.SyncPid{
				Pid:int32(pid),
			}
			p.SendMsg(201,proto_msgold)
		}
	}
	/*	proto_msg:=&pb.BroadCast{
			Pid:p.Pid,
			Tp:4,
			Data:&pb.BroadCast_P{
				&pb.Position{
					X:p.X,
					Y:p.Y,
					Z:p.Z,
					V:p.V,
				},
			},
		}
		players:=p.GetSurroundingPlayers()
		proto_players := make([]*pb.Player, 0, len(players))
		for _,v:=range players{
			v.SendMsg(200,proto_msg)
			proto_players = append(proto_players,
				&pb.Player{
					Pid: v.Pid,
					P: &pb.Position{
						X: v.X,
						Y: v.Y,
						Z: v.Z,
						V: v.V},
				})
		}
		proto_msg1:=&pb.SyncPlayers{
			Ps:proto_players,
		}
		p.SendMsg(202,proto_msg1)*/
}
