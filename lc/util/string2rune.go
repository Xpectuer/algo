package util

func string2rune(s string) []rune {
	if len(s) == 0 {
		return []rune{}
	}
	ru := make([]rune, len(s))
	for i, r := range s {
		ru[i] = r
	}
	return ru
}
