package main

func middleNode_20220803(head *ListNode) *ListNode {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func middleNode_20220803_2(head *ListNode) *ListNode {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	curr = head
	for i := 0; i < length/2; i++ {
		curr = curr.Next
	}
	return curr
}
