package LinkedList

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	sDummy := &ListNode{-1, nil}
	lDummy := &ListNode{-1, nil}

	sPre, lPre, curr := sDummy, lDummy, head
	for curr != nil {

		//fmt.Println("start looping")
		if curr.Val < x {
			sPre.Next = curr
			sPre = sPre.Next
		} else if curr.Val >= x {
			lPre.Next = curr
			lPre = lPre.Next
		}
		curr = curr.Next
	}
	lPre.Next = nil
	sPre.Next = lDummy.Next
	return sDummy.Next
}
