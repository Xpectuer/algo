package LinkedList

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	pre := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			pre.Next = l1
			pre = l1
			l1 = l1.Next
		} else {
			pre.Next = l2
			pre = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		pre.Next = l1
	} else if l2 != nil {
		pre.Next = l2
	}

	// rest of the list
	return dummy.Next
}
