package wordSearchMy

func findWords(board [][]byte, words []string) []string {
	trie := NewTrie()
	m, n := len(board), len(board[0])
	for _, w := range words {
		trie.Insert(w)
	}

	res := []string{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if trie.root.child[board[i][j]] != nil {
				trie.Search(i, j, &board, &res, trie.root)
			}

		}
	}
	return res
}

type trieNode struct {
	word  string
	child map[byte]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{"$", map[byte]*trieNode{}}
}

// Trie is the exported struct of the data struct
type Trie struct {
	root *trieNode
}

/** Initialize your data structure here. */
func NewTrie() *Trie {
	return &Trie{root: newTrieNode()}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this.root
	for _, w := range word {
		bw := byte(w)
		if _, ok := curr.child[bw]; !ok {
			curr.child[bw] = newTrieNode()
			if curr.word == "" {
				curr.word = "$"
			}

		}
		curr = curr.child[bw]
	}
	curr.word = word

}

func (this *Trie) IsLeaf(curr *trieNode) bool {
	for _, v := range curr.child {
		if v != nil {
			return false
		}
	}
	return true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(i, j int, board *[][]byte, res *[]string, curr *trieNode) {
	if i < 0 || i >= len(*board) || j < 0 || j >= len((*board)[0]) || (*board)[i][j] == '#' || curr.child[(*board)[i][j]] == nil {
		return
	}

	letter := (*board)[i][j]
	//parent := curr
	curr = curr.child[letter]

	//fmt.Println(string(letter), curr.word)
	// check match
	if curr.word != "$" {
		// fmt.Println("get word")
		(*res) = append((*res), curr.word)
		curr.word = "$"
	}

	// visited
	(*board)[i][j] = '#'
	movx, movy := [4]int{1, 0, -1, 0}, [4]int{0, 1, 0, -1}
	for k := 0; k < 4; k++ {
		//fmt.Println(i)
		this.Search(i+movx[k], j+movy[k], board, res, curr)
	}

	(*board)[i][j] = letter
	//curr = parent

	if this.IsLeaf(curr) {
		curr.child[letter] = nil
	}
}
