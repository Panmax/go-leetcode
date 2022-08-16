package main

func mergeTwoLists_20220816(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else {
		if list1.Val < list2.Val {
			list1.Next = mergeTwoLists_20220816(list1.Next, list2)
			return list1
		} else {
			list2.Next = mergeTwoLists_20220816(list1, list2.Next)
			return list2
		}
	}
}
