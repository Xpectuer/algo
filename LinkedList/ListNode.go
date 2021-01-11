package LinkedList

import "fmt"

// LC.86 LinkedList partition

// ListNode is a struct that defined by LC
type ListNode struct {
	Val  int
	Next *ListNode
}

func getLinkedList(arr []int) *ListNode {
	dummy := &ListNode{-1, nil}
	nxt := dummy
	for _, e := range arr {
		newNode := &ListNode{e, nil}
		nxt.Next = newNode
		nxt = newNode
	}
	return dummy.Next
}

func printLinkedList(head *ListNode) {
	if head == nil {
		return
	}
	if head.Next != nil {
		fmt.Printf("%d->", head.Val)
	} else {
		fmt.Printf("%d\n", head.Val)
	}

	printLinkedList(head.Next)
}
