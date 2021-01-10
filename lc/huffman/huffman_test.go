package huffman

import "testing"

func TestHuffman(t *testing.T) {
	var forest []*huffmanNode
	weight := []int{2, 3, 4, 5}
	char := []byte{'a', 'l', 'e', 'x'}
	for i := 0; i < len(weight); i++ {
		forest = append(forest, &huffmanNode{weight[i], char[i], nil, nil})
	}
	// init end
	t.Log("Init end... ")

	root := constructHuffman(forest)
	res := map[byte]int{}
	GetCode(root, 0, &res)
	t.Log("res: ", res)
	t.Log("End of test...")
}
