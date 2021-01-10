package wordSearch

import (
	"testing"
)

func TestWordSearch(t *testing.T) {

	board := [][]string{{"a", "b"}, {"a", "a"}}
	words := []string{"aba", "baa", "bab", "aaab", "aaa", "aaaa", "aaba"}
	nboard := make([][]byte, len(board))
	for i := range nboard {
		nboard[i] = make([]byte, len(board[0]))
	}
	for i := range board {
		for j := range board[i] {
			for _, c := range board[i][j] {
				nboard[i][j] = byte(c)
			}

		}
	}

	res := findWords(nboard, words)
	t.Log(res)
	t.Log("end of test")
}
