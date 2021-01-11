package stack

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {

	var (
		ans       string
		backtrack func(i int, s string) (int, string)
	)
	backtrack = func(i int, s string) (int, string) {
		res, time := "", 0
		for ; i < len(s); i++ {
			if s[i] >= '0' && s[i] <= '9' {
				tmp, _ := strconv.Atoi(string(s[i]))
				time = time*10 + tmp
				fmt.Println(time)
			} else if s[i] == '[' {
				tmp := ""
				i, tmp = backtrack(i+1, s)

				for k := 0; k < time; k++ {
					res += tmp
					//fmt.Println(res)
				}
				time = 0
			} else if s[i] == ']' {
				return i, res
			} else {
				res += string(s[i])
			}

		}
		return i, res

	}
	_, ans = backtrack(0, s)

	return ans
}

// Using asistant satck
func decodeStringStk(s string) string {
	type tup struct {
		time int
		res  string
	}
	var (
		stack []tup
	)
	stack, res, time := []tup{}, "", 0
	for _, c := range s {
		if c == '[' {
			stack = append(stack, tup{time: time, res: res})
			res, time = "", 0
		} else if c == ']' {
			t := stack[len(stack)-1]
			// pop
			stack = stack[:len(stack)-1]
			fmt.Println(t.time)
			for k := 0; k < t.time; k++ {
				t.res += res
			}
			res = t.res
		} else if c >= '0' && c <= '9' {
			time = time*10 + int(c-'0')
		} else {
			res += string(c)
		}
	}
	return res
}
