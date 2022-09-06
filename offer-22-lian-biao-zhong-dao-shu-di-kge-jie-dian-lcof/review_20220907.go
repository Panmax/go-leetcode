package main

func getKthFromEnd_20220907(head *ListNode, k int) *ListNode {
	slow := head
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}
