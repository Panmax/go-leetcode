package main

func detectCycle_20220810(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			ptr := head
			for slow != ptr {
				slow = slow.Next
				ptr = ptr.Next
				return slow
			}
		}
	}
	return nil
}
