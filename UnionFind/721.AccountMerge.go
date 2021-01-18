package disjoinSet

import "sort"

// TLE , just use integer as key to pass the TL
// DFS solution, see 721DFS.go
func accountsMerge(accounts [][]string) [][]string {
	// to initiate an /email->name/ map
	names := map[string]string{}
	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			names[account[i]] = account[0]
		}
	}
	//fmt.Println("names: ", names)

	size := len(names)
	parent := make(map[string]string, size)
	rank := make(map[string]int, size)
	for i := range rank {
		rank[i] = 1
	}
	// k: name v: email

	var Find = func(key string) string {
		if _, ok := parent[key]; !ok {
			parent[key] = key
			return key
		}
		curr := key
		for curr != parent[curr] {
			// route compression
			parent[curr] = parent[parent[curr]]
		}

		return curr
	}

	var Union = func(x, y string) {
		rootX, rootY := Find(x), Find(y)
		if rootX == rootY {
			return
		}
		// notice that X is as a root ,otherwise a bug occur
		// union by rank
		if rank[rootX] < rank[rootY] {
			parent[rootX] = rootY
			rank[rootY] += rank[rootX]
		} else {
			parent[rootY] = rootX
			rank[rootX] += rank[rootY]
		}
		// Find(x)
		// Find(y)
		// route compression after union to debug

		return
	}

	// TODO:[FIXED] johnsmith not in the set expect in the set
	// using rank
	for _, account := range accounts {
		// problem
		for i := 1; i < len(account); i++ {
			// fmt.Println(account[i])
			Union(account[1], account[i])
		}
	}
	//fmt.Println("parent:", parent)
	// fill results for each set
	resMap := make(map[string][]string, len(names))
	for k, v := range parent {
		resMap[v] = append(resMap[v], k)
	}
	// parse the results
	ress := make([][]string, 0)
	for k, v := range resMap {
		res := make([]string, 0)
		res = append(res, names[k])
		res = append(res, v...)
		sort.Strings(res)
		ress = append(ress, res)
	}

	return ress
}
