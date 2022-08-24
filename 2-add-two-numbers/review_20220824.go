package main

func addTwoNumbers_20220824(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	dummy := &ListNode{}
	head := dummy

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
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

	// 注意这里最后判断下 carry
	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
