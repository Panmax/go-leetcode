package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{Val: -1}
	next := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			next.Next = list1
			list1 = list1.Next
		} else {
			next.Next = list2
			list2 = list2.Next
		}
		next = next.Next
	}

	if list1 != nil {
		next.Next = list1
	} else {
		next.Next = list2
	}

	return dummy.Next
}

func main() {

}
