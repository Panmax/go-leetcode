package main

func mergeKLists_20221014(lists []*ListNode) *ListNode {
	return merge_20221014(lists, 0, len(lists)-1)
}

func merge_20221014(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	// 需要判断 left > right 的情况
	if left > right {
		return nil
	}

	mid := (right + left) / 2
	return mergeTwoList_20221014(merge_20221014(lists, left, mid), merge_20221014(lists, mid+1, right))
}

func mergeTwoList_20221014(list1, list2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
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

	return head.Next
}
