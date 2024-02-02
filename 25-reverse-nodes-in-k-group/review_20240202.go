package main

func reverseKGroup_20240202(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair

	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		next := tail.Next
		head, tail = reverse_20240202(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = next
	}
	return hair.Next
}

func reverse_20240202(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}
