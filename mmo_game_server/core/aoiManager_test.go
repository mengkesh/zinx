package core

import (
	"testing"
	"fmt"
)

func TestAoiManager(t *testing.T){
	aoi:=NewAoiManager(0,250,0,250,5,5)
	fmt.Println(aoi)
}
func TestAoiManagerSurround(t *testing.T){
	aoi:=NewAoiManager(0,250,0,250,5,5)
	for gid,_:=range  aoi.grids{
		grids:=aoi.GetSurroundGridsByGid(gid)
		fmt.Println("gid : ", gid, " grids num = ", len(grids))
		gIDs := make([]int, 0, len(grids))
		for _,v:=range grids{
			gIDs=append(gIDs,v.GID)
		}
		fmt.Println("grids IDs are ", gIDs)
	}
	fmt.Println("======================")
	aoi.AddPidToGrid(1,2)
	aoi.AddPidToGrid(2,7)
	players:=aoi.GetSurroundPIDsByPos(175,68)
	fmt.Println(players)

}