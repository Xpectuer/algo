package playground

import "testing"

type TrieNode struct {
	word     string
	Children []*TrieNode
}

func TestPG(t *testing.T) {
	tn := &TrieNode{word: "$", Children: make([]*TrieNode, 26)}
	t.Log(tn.Children)

	hashmap := map[int]*TrieNode{1: nil}
	_, ok := hashmap[1]
	t.Log(ok)
}
