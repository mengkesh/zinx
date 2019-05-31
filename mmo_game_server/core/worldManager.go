package core

import "sync"

const (
	AOI_MIN_X int = 85
	AOI_MAX_X int = 410
	AOI_CNTS_X int = 10
	AOI_MIN_Y int = 75
	AOI_MAX_Y int = 400
	AOI_CNTS_Y int = 20
)
type WorldManager struct {
	Players map[int32]*player
	pLock sync.RWMutex
	AoiMgr *AoiManager
}
var WorldMgrObj *WorldManager

func init (){
	WorldMgrObj=NewWorldManager()
}
func NewWorldManager()*WorldManager{
	w:=&WorldManager{
		Players:make(map[int32]*player),
		AoiMgr:NewAoiManager(AOI_MIN_X,AOI_MAX_X,AOI_MIN_Y,AOI_MAX_Y,AOI_CNTS_X,AOI_CNTS_Y),
	}
	return w
}
func(wm *WorldManager)AddPlayer(player *player){
	wm.pLock.Lock()
	defer wm.pLock.Unlock()
	wm.Players[player.Pid]=player
	wm.AoiMgr.AddToGridByPos(int(player.Pid) ,player.X,player.Z)
}
func(wm *WorldManager)RemovePlayerByPid(pid int32){
	wm.pLock.Lock()
	defer wm.pLock.Unlock()
	wm.AoiMgr.RemoveFromGridbyPos(int(pid),wm.Players[pid].X,wm.Players[pid].Z)
	delete(wm.Players,pid)
}
func(wm *WorldManager)GetPlayerByPid(pid int32)*player{
	wm.pLock.RLock()
	p:=wm.Players[pid]
	wm.pLock.RUnlock()
	return p
}
func(wm *WorldManager)GetAllPlayers()[]*player{
	wm.pLock.Lock()
	defer wm.pLock.Unlock()
	var p []*player
	for _,v:=range wm.Players{
		p=append(p,v)
	}
	return p
}