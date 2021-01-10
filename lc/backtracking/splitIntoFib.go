package backtracking

import "math"

func splitIntoFibonacci(s string) (F []int) {
	n := len(s)
	var backtrack func(index, sum, prev int) bool
	backtrack = func(index, sum, prev int) bool {
		if index == n {
			return len(F) >= 3
		}

		cur := 0
		for i := index; i < n; i++ {
			// check if started with 0
			if i > index && s[index] == '0' {
				break
			}

			cur = cur*10 + int(s[i]-'0')
			// validate the uint_32
			if cur > math.MaxInt32 {
				break
			}

			// f[i] + f[i+1] = f[i+2]
			if len(F) >= 2 {
				if cur < sum {
					continue
				}
				if cur > sum {
					break
				}
			}

			// valid cur
			F = append(F, cur)
			if backtrack(i+1, prev+cur, cur) {
				return true
			}
			// restore the F
			F = F[:len(F)-1]
		}
		return false
	}
	backtrack(0, 0, 0)
	return
}
