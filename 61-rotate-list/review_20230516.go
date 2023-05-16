package main

func rotateRight_20230516(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	cur := head
	for cur.Next != nil {
		n++
		cur = cur.Next
	}
	add := n - k%n
	if add == n {
		return head
	}
	cur.Next = head
	for add > 0 {
		cur = cur.Next
		add--
	}
	node := cur.Next
	cur.Next = nil
	return node
}
