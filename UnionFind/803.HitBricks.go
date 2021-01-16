package disjoinSet

import "fmt"

// leetcode 803. difficulty: HARD

//grid:=[][]int{{1,0,0,0},
//				{1,1,1,0}}
//	hits := [][]int{{1,0}}

func hitBricks(grid [][]int, hits [][]int) []int {
	// inverse thought:
	// we fill bricks been hit back to the grid
	rows := len(grid)
	cols := len(grid[0])
	// slice is a pointer variable so this will occur shallow copy
	//copyArr := grid
	copyArr := make([][]int,len(grid))
	for i := range copyArr {
		copyArr[i] = make([]int, len(grid[i]))
	}
	for i := range grid {
		for j := range grid[i] {
			copyArr[i][j] = grid[i][j]
		}
	}


	for _,hit := range hits {
		copyArr[hit[0]][hit[1]] = 0
	}


	fmt.Printf("copy:\n%v\n%v\n",copyArr[0],copyArr[1])

	// funcs
	var getIndex = func(x, y int) int {
		return x * cols + y
	}
	var inArea = func(x, y int) bool {
		return x < rows && x>=0 && y<cols && y >=0
	}
	// size idx as a root
	size := cols * rows
	// 并查集的底层
	parent := make([]int,size+1)
	// 并查集每个元素的秩
	ufsize := make([]int,size+1)

	// initiating
	// -----------------------------------
	for i := 0;i<=size;i++ {
		parent[i] = i
		ufsize[i] = 1
	}

	var Find = func(x int) int {
		root := x
		for root != parent[root] {
			root = parent[root]
		}
		// path compression
		parent[x] = root
		return root

	}

	var Union = func(x, y int) {
		rootX, rootY := Find(x), Find(y)
		if rootX == rootY {
			return
		}
		if ufsize[rootX] < ufsize[rootY] {
			parent[rootX] = rootY
			ufsize[rootY] += ufsize[rootX]
		} else {
			parent[rootY] = rootX
			ufsize[rootX] += ufsize[rootY]
		}


		return
	}

	var getSize = func(x int) int {
		root:= Find(x)
		return ufsize[root]
	}

	// -----------------------------------

	// step1. handle the roof
	for j:=0; j< cols ; j++ {
		if copyArr[0][j] == 1 {
			Union(j,size)
		}
	}

	// step2. handle the row No.2 - end
	for i:=1;i<rows;i++ {
		for j:=0;j<cols;j++ {
			if copyArr[i][j] == 1 {
				if copyArr[i-1][j] == 1 {
					Union(getIndex(i-1,j),getIndex(i,j))
				}
				if j>0 && copyArr[i][j-1] == 1 {
					Union(getIndex(i,j-1),getIndex(i,j))
				}
			}
		}
	}
	fmt.Println("parent: ",parent)

	hitsLen := len(hits)
	res := make([]int,hitsLen)
	// step3. traverse the hits array,
	// fill the bricks back to the copy array
	for i:= hitsLen-1;i>=0;i-- {
		// fill the bricks
		x := hits[i][0]
		y := hits[i][1]
		// original empty
		if grid[x][y] == 0 {
			continue
		}

		// original ufsize of root node
		origin := getSize(size)
		//fmt.Println(origin)

		// if there is a roof
		if x==0 {
			Union(y, size)
		}
		var (
			MVX = [4]int{0,1,0,-1}
			MVY = [4]int{1,0,-1,0}
		)

		for k:=0;k<4;k++ {
			newX := x + MVX[k]
			newY := y + MVY[k]

			if inArea(newX,newY)&& copyArr[newX][newY] == 1 {
				Union(getIndex(x,y),getIndex(newX,newY))
			}
		}

		// current ufsize of root node
		current := getSize(size)

		if current - origin-1 < 0 {
			res[i] =0
		} else {
			res[i] = current - origin -1
		}
		// fill the brick
		copyArr[x][y] = 1
 	}

	return res
}
