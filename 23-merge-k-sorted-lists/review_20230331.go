package main

func mergeKLists_20230331(lists []*ListNode) *ListNode {
	return merge_20230331(lists, 0, len(lists)-1)
}

func merge_20230331(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoList_20230331(merge_20230331(lists, l, mid), merge_20230331(lists, mid+1, r))
}

func mergeTwoList_20230331(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	if list1 != nil {
		tail.Next = list1
	} else if list2 != nil {
		tail.Next = list2
	}
	return dummy.Next
}
