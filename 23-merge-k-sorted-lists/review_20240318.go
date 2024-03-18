package main

func mergeKLists_20240318(lists []*ListNode) *ListNode {
	return merge_20240318(lists, 0, len(lists)-1)
}

func merge_20240318(lists []*ListNode, l int, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := (l + r) / 2
	return mergeTwoList_20240318(merge_20240318(lists, l, mid), merge_20240318(lists, mid+1, r))
}

func mergeTwoList_20240318(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}

	return dummy.Next
}
