package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow := &ListNode{
		Next: head,
	}
	fast := slow
	index := 0
	for index != k {
		fast = fast.Next
		index++
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow.Next
}
