package main

func removeNthFromEnd_20220829(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	i := 0
	for fast != nil && i != n {
		fast = fast.Next
		i++
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
