package main

func addTwoNumbers_20220921(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	var carry int
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		head.Next = &ListNode{Val: sum}
		head = head.Next
	}
	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
