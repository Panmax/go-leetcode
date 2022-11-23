package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	cur := head
	for cur.Next != nil {
		cur = cur.Next
		n++
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
	ret := cur.Next
	cur.Next = nil
	return ret
}
