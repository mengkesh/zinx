package core

import "fmt"

type AoiManager struct {
	MinX  int
	MaxX  int
	MinY  int
	MaxY  int
	CntsX int
	CntsY int
	grids map[int]*Grid
}

func (this *AoiManager) Getwidth() int {
	return (this.MaxX - this.MinX) / this.CntsX
}
func (this *AoiManager) Getheight() int {
	return (this.MaxY - this.MinY) / this.CntsY
}
func NewAoiManager(minx, maxx, miny, maxy, cntsx, cntsy int)*AoiManager{
	aoimgr := &AoiManager{
		MinX:  minx,
		MaxX:  maxx,
		MinY:  miny,
		MaxY:  maxy,
		CntsX: cntsx,
		CntsY: cntsy,
		grids: make(map[int]*Grid),
	}
	for y := 0; y < cntsy; y++ {
		for x := 0; x < cntsx; x++ {
			gid := y*cntsx + x
			aoimgr.grids[gid] = NewGrid(gid,
				x*aoimgr.Getwidth(),
				(x+1)*aoimgr.Getwidth(),
				y*aoimgr.Getheight(),
				(y+1)*aoimgr.Getheight(),
			)
		}
	}
	return aoimgr
}
func(this *AoiManager)String()string{
	s:= fmt.Sprintf("MinX:%d,MaxX:%d,MinY:%d,MaxY:%d,CntsX:%d,CntsY:%d\n",
		this.MinX,this.MaxX,this.MinY,this.MaxY,this.CntsX,this.CntsY)
	for _,grid:=range this.grids{
		s+=fmt.Sprintln(grid)
	}

	return s
}
func(this *AoiManager)AddPidToGrid(pid,gid int){
	this.grids[gid].Add(pid,nil)
}
func(this *AoiManager)RemovePidFromGrid(pid,gid int) {
	this.grids[gid].Remove(pid)
}
func(this *AoiManager)GetPidsByGid(gid int)[]int{
	playersid:=this.grids[gid].GetPlayerID()
	return playersid
}
func(this *AoiManager)GetSurroundGridsByGid(gID int) (grids []*Grid){
	if _,ok:=this.grids[gID];!ok{
		return
	}
	grids=append(grids,this.grids[gID])
	idx:=gID%this.CntsX
	if idx>0{
		grids=append(grids,this.grids[gID-1])
	}
	if  idx<this.CntsX-1{
			grids=append(grids,this.grids[gID+1])
	}
	gidsX:=make([]int,0,len(grids))
	for _,v:=range grids{
		gidsX=append(gidsX,v.GID)
	}
	for _,gid:=range gidsX{
		idy:=gid/this.CntsX
		if idy>0{
			grids=append(grids,this.grids[gid-this.CntsX])
		}
		if idy<this.CntsY-1{
			grids=append(grids,this.grids[gid+this.CntsX])
		}
	}
	return
}
func(this *AoiManager)GetGidByPos(x,y float32)int{
	if x<0 || x>float32(this.MaxX){
		return  -1
	}
	if y<0 ||y>float32(this.MaxY){
		return  -1
	}
	idx:=(int(x)-this.MinX)/this.Getwidth()
	idy:=(int(y)-this.MinY)/this.Getheight()
	return idy*this.CntsX+idx
}
func(this *AoiManager)GetSurroundPIDsByPos(x,y float32)[]int{
	gid:=this.GetGidByPos(x,y)
	fmt.Println(gid)
	grids:=this.GetSurroundGridsByGid(gid)
	var allpids []int
	for _,v:=range grids{
		allpids=append(allpids,v.GetPlayerID()...)
	}
	return allpids
}
func(this *AoiManager)AddToGridByPos(pID int,x,y float32){
	gid:=this.GetGidByPos(x,y)
	this.grids[gid].Add(pID,nil)
}
func(this *AoiManager)RemoveFromGridbyPos(pID int,x,y float32){
	gid:=this.GetGidByPos(x,y)
	this.grids[gid].Remove(pID)
}