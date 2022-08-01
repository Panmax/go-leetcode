package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	rList := reverseList(list)
	fmt.Println(rList)
}
