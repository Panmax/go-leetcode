package main

func trainingPlan_20240226(head *ListNode, cnt int) *ListNode {
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
