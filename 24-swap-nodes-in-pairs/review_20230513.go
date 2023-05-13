package main

func swapPairs_20230513(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newNode := swapPairs(next.Next)
	next.Next = head
	head.Next = newNode

	// return next
	return next
}
