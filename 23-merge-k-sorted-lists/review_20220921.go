package main

func mergeKLists_20220921(lists []*ListNode) *ListNode {
	return merge_20220921(lists, 0, len(lists)-1)
}

func merge_20220921(lists []*ListNode, l, r int) *ListNode {
	// 这里的逻辑需要再熟悉下
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoList_20220921(merge_20220921(lists, l, mid), merge_20220921(lists, mid+1, r))
}

func mergeTwoList_20220921(list1, list2 *ListNode) *ListNode {
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
