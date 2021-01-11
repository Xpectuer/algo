package trie

type trieNode struct {
	end   bool
	child map[rune]*trieNode
}

func newTrieNode() *trieNode {
	return &trieNode{false, map[rune]*trieNode{}}
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
