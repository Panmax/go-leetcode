package main

func getIntersectionNode_20240227(headA, headB *ListNode) *ListNode {
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
