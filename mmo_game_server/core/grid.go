package core

import (
	"sync"
	"fmt"
)

type Grid struct {
	GID       int
	MinX      int
	MaxX      int
	MinY      int
	MaxY      int
	playerInfos map[int]interface{}
	pIDLock   sync.RWMutex
}

func NewGrid(gid, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerInfos: make(map[int]interface{}),
	}

}
func(this *Grid)Add(playerid int,player interface{}){
	this.pIDLock.Lock()
	defer this.pIDLock.Unlock()
	this.playerInfos[playerid]=player
}
func(this *Grid)GetPlayerID()[]int{
	this.pIDLock.RLock()
	defer this.pIDLock.RUnlock()
	var playerids []int
	for k,_:=range this.playerInfos{
		playerids=append(playerids,k)
	}
	return playerids
}
func (this *Grid)Remove(playerid int){
	this.pIDLock.Lock()
	defer this.pIDLock.Unlock()
	delete(this.playerInfos,playerid)
}
func(this *Grid)String()string{
	return fmt.Sprintf("Grid id:%d,minx:%d,maxx:%d,miny:%d,maxy:%d,players:%v\n" ,
		this.GID,this.MinX,this.MaxX,this.MinY,this.MaxY,this.playerInfos)
}