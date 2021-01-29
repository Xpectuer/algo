package search

import "sort"

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 */

/*
 search使用二分法进行查找，Search()方法回使用“二分查找”算法来搜索某指定切片[0:n]，
 并返回能够使f(i)=true的最小的i（0<=i<n）值，并且会假定，
 如果f(i)=true，则f(i+1)=true，
 即对于切片[0:n]，i之前的切片元素会使f()函数返回false，i及i之后的元素会使f()函数返回true。
 但是，当在切片中无法找到时f(i)=true的i时（
 此时切片元素都不能使f()函数返回true），Search()方法会返回n（而不是返回-1）。
*/
type pair struct{ x, y int }

var dirs = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func minimumEffortPath(heights [][]int) int {
	n, m := len(heights), len(heights[0])
	return sort.Search(1e6, func(maxHeightDiff int) bool {
		vis := make([][]bool, n)
		for i := range vis {
			vis[i] = make([]bool, m)
		}
		vis[0][0] = true
		queue := []pair{{}}
		for len(queue) > 0 {
			// BFS
			p := queue[0]
			queue = queue[1:]
			if p.x == n-1 && p.y == m-1 {
				return true
			}
			for _, d := range dirs {
				x, y := p.x+d.x, p.y+d.y
				// ...<= maxHeightDiff / 如果大于，说明无法登上
				if 0 <= x && x < n && 0 <= y && y < m && !vis[x][y] && abs(heights[x][y]-heights[p.x][p.y]) <= maxHeightDiff {
					vis[x][y] = true
					queue = append(queue, pair{x, y})
				}
			}
		}
		return false
	})
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/path-with-minimum-effort/solution/zui-xiao-ti-li-xiao-hao-lu-jing-by-leetc-3q2j/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
