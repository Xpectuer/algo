package trieNew

type TrieNode struct {
  end bool
	word     string
	Children []*TrieNode
}

type Trie struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() Trie {
	root := &TrieNode{word: "$", Children: make([]*TrieNode, 26),end:false}
	return Trie{root: root}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}
	node := this.root
	for _, e := range word {
		if node.Children[e-'a'] == nil {
			node.Children[e-'a'] = &TrieNode{word: '$', Children: make([]*TrieNode, 26)}
		}
		node = node.Children[e-'a']
	}
	// add last with word
	node.word = word
  node.end = true

}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return this.find(word, true)
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return this.find(word, false)
}

func (this *Trie) find(word string, exact bool) bool {
  node := this.root
  for _, e:= word {
    if node.Children[e-'a']!=nil {
      return false
    } else {
      node = node.Children[e-'a']
    }
  }
  if exact {
    return node.end
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
