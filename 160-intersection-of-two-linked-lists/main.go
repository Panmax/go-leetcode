package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		m[tmp] = true
	}
	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if m[tmp] {
			return tmp
		}
	}
	return nil
}
