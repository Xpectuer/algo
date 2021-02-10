package slideWindow

//import (
//    "fmt"
//    )
func characterReplacement(s string, k int) int {
    n := len(s)
    if n <= 1 {return n}
    left, right := 0, 0
    feq := make(map[byte]int,52)
    maxCnt, res := 0 , 0

    for right < n {
        feq[s[right]]++
        maxCnt = max(maxCnt,feq[s[right]])
        right++
        if maxCnt + k < right - left {
            // k 不够用
            feq[s[left]]--
            left++
        }
        res = max(res, right - left)
    }
    return res
}

func max(a,b int) int {
    if a>b {
        return a
    }
    return b
}

