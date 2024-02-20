package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func trainingPlan(head *ListNode, cnt int) *ListNode {
	slow, fast := head, head
	for i := 0; i < cnt; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
