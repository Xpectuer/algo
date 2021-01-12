package topologicalSort

// leetcode 1203
func topSort(graph [][]int, deg, items []int) []int {
	q, orders := []int{}, []int{}
	for _, it := range items {
		if deg[it] == 0 {
			q = append(q, it)
		}
	}
	for len(q) > 0 {
		from := q[0]
		q = q[1:]
		orders = append(orders, from)
		for _, to := range graph[from] {
			deg[to]--
			if deg[to] == 0 {
				q = append(q, to)
			}
		}
	}
	return orders
}

func sortItems(n, m int, group []int, beforeItems [][]int) (ans []int) {
	groupItems := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			// allocate for items without a group
			group[i] = m + i
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	groupGraph := make([][]int, m+n)
	groupDegree := make([]int, m+n)
	itemGraph := make([][]int, n)
	itemDegree := make([]int, n)

	// init degreeTable
	for cur, items := range beforeItems {
		currGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != currGroupID { // intergroup relationship
				groupGraph[preGroupID] = append(groupGraph[preGroupID], currGroupID)
				groupDegree[currGroupID]++
			} else {
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	// topSort inter group
	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}
	groupOrders := topSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	// in-group topSort
	for _, groupID := range groupOrders {
		items := groupItems[groupID]
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return

}
