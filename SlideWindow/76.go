package slideWindow

import "fmt"

const INTMAX = 1 << 31

func m_minWindow(s string, t string) string {
	dict := map[byte]int{}
	for _, e := range t {
		dict[byte(e)]++
	}

	// hamming distance between the substring and t
	counter := len(t)
	// length of substring
	length := INTMAX
	// pointers
	begin, end, head := 0, 0, 0
	for end < len(s) {
		//fmt.Println("outer")

		
		if dict[s[end]] > 0 {
			counter--
		}
		dict[s[end]]--
		end++
		
		// find min window iteratively
		for counter == 0 {
			//fmt.Println("Inner")
			if end-begin < length {
				length = end - begin
				head = begin
			}

			// left bound ->
			
			if dict[s[begin]] == 0 {
				counter++
			}
			dict[s[begin]]++
			begin++
			
		}
	}

	if length == INTMAX {
		return ""
	}
	return s[head:head+length]
}
