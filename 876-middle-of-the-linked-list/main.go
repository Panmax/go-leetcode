package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 给定一个头结点为 head 的非空单链表，返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。
func middleNode_slowFast(head *ListNode) *ListNode {
	slow := head
	fast := head
	for slow != nil && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func middleNode_singlePoint(head *ListNode) *ListNode {
	length := 0
	curr := head
	for curr != nil {
		length++
		curr = curr.Next
	}
	curr = head
	for i := 0; i < length/2; i++ {
		curr = curr.Next
	}
	return curr
}

func middleNode(head *ListNode) *ListNode {
	nodeMap := make(map[int]*ListNode)
	length := 0
	for head != nil {
		nodeMap[length] = head
		length++
		head = head.Next
	}
	middle := length / 2
	return nodeMap[middle]
}
