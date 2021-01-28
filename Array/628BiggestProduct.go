package array

/*
 * @Author: Alex
 * @Date: 2021-01-20 15:53:46
 * @LastEditTime: 2021-01-20 16:44:56
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /algo/Array/628BiggestProduct.go
 */

// What if we need maximum Product of N numbers ?
// Dynamic Programming
func maximumProduct(nums []int) int {
	const (
		MAX = (1 << 32) - 1
		MIN = -MAX
	)
	// cap 3
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, 4)
	}
	dp[0][0], dp[1][0] = 1, 1
	for j := 3; j > 0; j-- {
		dp[0][j] = MIN
		dp[1][j] = MAX
	}

	for i := 0; i < len(nums); i++ {
		for j := 3; j > 0; j-- {
			if dp[0][j-1] == MIN {
				continue
			}
			dp[0][j] = max(dp[0][j], max(dp[0][j-1]*nums[i], dp[1][j-1]*nums[i]))
			dp[1][j] = min(dp[1][j], min(dp[0][j-1]*nums[i], dp[1][j-1]*nums[i]))
		}

	}
	return dp[0][3]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
