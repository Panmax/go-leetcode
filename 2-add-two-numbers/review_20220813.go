package main

func addTwoNumbers_20220813(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	carry := 0

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
		carry, sum = sum/10, sum%10

		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: sum,
			}
			tail = tail.Next
		}
	}

	if carry != 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}

	return head
}
