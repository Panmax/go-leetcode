package main

func hasCycle_20230302(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
