package main

func rotateRight_20221127(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n := 1
	curr := head
	for curr.Next != nil {
		n++
		curr = curr.Next
	}
	add := n - k%n
	if add == n {
		return head
	}
	curr.Next = head
	for add > 0 {
		curr = curr.Next
		add--
	}
	ret := curr.Next
	curr.Next = nil
	return ret
}
