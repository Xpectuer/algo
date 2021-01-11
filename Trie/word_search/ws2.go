package wordSearch

func findWords(board [][]byte, words []string) []string {
	// 1. construct the Trie
	trie := Constructor()

	for _, w := range words {
		trie.Insert(w)
	}
	// parameters
	// row, col := len(board), len(board[0])
	res := make([]string, 0)

	var backtrack func(i int, j int, parent *trieNode)
	backtrack = func(i int, j int, parent *trieNode) {
		l := board[i][j]
		curr := parent.child[rune(board[i][j])]

		if curr.word != "$" {
			res = append(res, curr.word)
		}

		rm, cm := []int{0, 1, 0, -1}, []int{-1, 0, 1, 0}

		for k := range []int{0, 1, 2, 3} {
			if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
				continue
			}
			newrow := i + rm[k]
			newcol := j + cm[k]

			if _, ok := curr.child[rune(board[newrow][newcol])]; ok {
				backtrack(newrow, newcol, curr)
			}
			board[i][j] = l

		}

	}
	for i := range board {
		for j := range board[i] {
			if _, ok := trie.root.child[rune(board[i][j])]; ok {
				backtrack(i, j, trie.root)
			}
		}
	}
	return res
}

type trieNode struct {
	end   bool
	child map[rune]*trieNode
	word  string
}

func newTrieNode() *trieNode {
	return &trieNode{false, map[rune]*trieNode{}, "$"}
}

// Trie is the exported struct of the data struct
type Trie struct {
	root *trieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	t := Trie{}
	t.root = newTrieNode()

	return t
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this.root
	for _, w := range word {
		if _, ok := curr.child[w]; !ok {
			curr.child[w] = newTrieNode()
		}
		curr = curr.child[w]
	}
	curr.end = true
	curr.word = word

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.find(word, true)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.find(prefix, false)
}

func (this *Trie) find(word string, exactMatch bool) bool {
	curr := this.root
	for _, w := range word {
		if curr.child[w] == nil {
			return false
		} else {
			curr = curr.child[w]
		}

	}
	if exactMatch {
		return curr.end
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
