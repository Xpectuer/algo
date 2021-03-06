package tree

import (
	"fmt"
	"math"
	"testing"
)

func TestPathSum2(t *testing.T) {
	A := []int{-6, math.MinInt32, -3, -6, 0, -6, -5, 4, math.MinInt32, math.MinInt32, math.MinInt32, 1, 7}

	targetSum := -21
	_ = targetSum
	root := newTreeNode(&A)
	//preorder(root)
	res := pathSum(root, targetSum)
	t.Log(res)
}

func newTreeNode(serial *[]int) *TreeNode {

	if len(*serial) == 0 {
		return nil
	}
	if (*serial)[0] == math.MinInt32 {
		(*serial) = (*serial)[1:]
		return nil
	}

	Val := (*serial)[0]
	node := &TreeNode{Val: Val, Left: nil, Right: nil}
	(*serial) = (*serial)[1:]
	node.Left = newTreeNode(serial)
	node.Right = newTreeNode(serial)

	return node
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorder(root.Left)
	preorder(root.Right)
}
