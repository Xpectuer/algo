package disjoinSet

import "testing"

func Test803(t *testing.T) {
	grid:=[][]int{{1,0,0,0},{1,1,1,0}}
	hits := [][]int{{1,0}}
	t.Logf("grid: \n%v \n%v",grid[0],grid[1])
	t.Log("hits:",hits)
	//t.Log("result: ",hitBricks(grid,hits))
	arr := make([][]int,len(grid))
	for i :=range arr {arr[i]=make([]int,len(grid[0]))}
	copy(arr,grid)
	t.Log(arr)
}