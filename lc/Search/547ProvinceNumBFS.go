package search

func findCircleNumBFS(isConnected [][]int) int {
	// dfs
	n := len(isConnected)
	visited := make([]bool, n)
	q := make([]int, 0)
	cnt := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			cnt++
			q = append(q, i)
			visited[i] = true
			// BFS
			for len(q) != 0 {
				v := q[0]
				q = q[1:] // pop
				for j := 0; j < n; j++ {
					if isConnected[v][j] == 1 && !visited[j] {
						visited[j] = true
						q = append(q, j)

					}
				}
			}

		}

	}
	return cnt
}
