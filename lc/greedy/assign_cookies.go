package greedy

import "sort"

// leetcode 455.
func findContentChildren(g []int, s []int) int {
	// sort first
	m, n := len(g), len(s)
	sort.Ints(g)
	sort.Ints(s)

	count := 0

	for idx1, idx2 := 0, 0; idx1 < m; idx1++ {
		for idx2 < n {
			if g[idx1] <= s[idx2] {
				idx2++
				count++
				break

			} else {
				idx2++
			}
		}

	}

	return count

}
