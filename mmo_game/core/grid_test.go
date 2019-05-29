package core

import (
	"testing"
	"fmt"
)

func TestGrid(t *testing.T){
	player1:="player1"
	player2:="player2"
	grid:=NewGrid(1,0,4,3,6)
	grid.Add(1,player1)
	grid.Add(2,player2)
	fmt.Println(grid)
}