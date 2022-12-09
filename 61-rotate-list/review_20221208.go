package main

func rotateRight_20221208(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
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
