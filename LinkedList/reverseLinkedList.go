package LinkedList

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	ret := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return ret
}

func reverseItVer(head *ListNode) *ListNode {
	var (
		prev *ListNode
		curr *ListNode
	)
	prev, curr = nil, head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}
