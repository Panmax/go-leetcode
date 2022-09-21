package main

func mergeTwoLists_20220921(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	head := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}
	if list1 == nil {
		head.Next = list2
	} else if list2 == nil {
		head.Next = list1
	}

	return dummy.Next
}
