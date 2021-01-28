package slideWindow

/*
 * @Author: XPectuer
 * @LastEditor: XPectuer
 * @REF: Leetcode
 */

func minWindow(s string, t string) string {
	sLen, tLen := len(s), len(t)
	if sLen == 0 || tLen == 0 || sLen < tLen {
		return ""
	}

	winFreq, tFreq := make(map[byte]int, sLen*2), make(map[byte]int, tLen*2)
	for i := range t {
		tFreq[t[i]]++
	}
	// number of chars window contains in t
	distance := 0
	minLen := sLen + 1
	begin := 0

	left, right := 0, 0
	// [left,right) 16 26
	for right < sLen {
		// right bound slide right
		charRight := tFreq[s[right]]
		// charLeft := tFreq[s[left]]
		if charRight == 0 {
			right++
			continue
		}
		if winFreq[s[right]] < charRight {
			distance++
		}
		winFreq[s[right]]++
		right++
		for distance == tLen {
			if right-left < minLen {
				minLen = right - left
				begin = left
			}
			// char not in t
			if tFreq[s[left]] == 0 {
				left++
				continue
			}
			if winFreq[s[left]] == tFreq[s[left]] {
				distance--
			}
			winFreq[s[left]]--
			left++
		}

	}

	if minLen == sLen+1 {
		return ""
	}
	return s[begin : begin+minLen]
}
