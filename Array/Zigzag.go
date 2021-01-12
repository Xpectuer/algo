package array

import "strings"

// leetcode 6.
// ref: https://leetcode-cn.com/problems/zigzag-conversion/solution/zzi-xing-bian-huan-by-jyd/
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	// StringBuilder recommended
	res := make([]strings.Builder, numRows)

	// simulate the zigzag procedure
	i, flag := 0, -1
	for _, c := range s {

		res[i].WriteRune(c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	var sb strings.Builder
	sb.Grow(len(s))
	for _, s := range res {
		sb.WriteString(s.String())
	}
	return sb.String()
}
