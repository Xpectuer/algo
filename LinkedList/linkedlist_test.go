package LinkedList

import (
	"testing"
)

func TestPartition(t *testing.T) {
	arr := []int{1, 4, 3, 2, 5, 2}
	head := getLinkedList(arr)
	head = partition(head, 3)
	printLinkedList(head)
}

func TestReverse(t *testing.T) {
	arr := []int{1, 4, 3, 2, 5, 2}
	head := getLinkedList(arr)
	head = reverseItVer(head)
	printLinkedList(head)
}

func TestMerge(t *testing.T) {
	l1 := []int{1, 2, 4}
	l2 := []int{1, 3, 4}

	h1 := getLinkedList(l1)
	h2 := getLinkedList(l2)

	ret := mergeTwoLists(h1, h2)
	printLinkedList(ret)
}
