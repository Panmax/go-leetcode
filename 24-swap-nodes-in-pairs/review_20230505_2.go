package main

func swapPairs_20230505_2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	next := head.Next
	newNode := swapPairs(next.Next)
	next.Next = head
	head.Next = newNode

	return next
}
