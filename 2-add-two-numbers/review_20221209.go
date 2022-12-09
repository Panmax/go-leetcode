package main

func addTwoNumbers_20221209(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy
	var carry int
	for l1 != nil || l2 != nil {
		var num1, num2 int
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}
		sum := num1 + num2 + carry
		sum, carry = sum%10, sum/10
		head.Next = &ListNode{Val: sum}
		head = head.Next
	}
	if carry != 0 {
		head.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
