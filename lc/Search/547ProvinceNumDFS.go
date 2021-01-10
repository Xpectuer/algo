package search

func findCircleNum(isConnected [][]int) int {
	// dfs
	n := len(isConnected)
	visited := make([]bool, n)
	cnt := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			cnt++
			dfs(i, isConnected, visited)
		}
	}
	return cnt
}

func dfs(i int, isConnected [][]int, visited []bool) {
	visited[i] = true

	for j := 0; j < len(isConnected); j++ {
		if isConnected[i][j] == 1 && !visited[j] {
			dfs(j, isConnected, visited)
		}

	}
}
