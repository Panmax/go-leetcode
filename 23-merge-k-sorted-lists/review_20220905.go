package main

func mergeKLists_20220905(lists []*ListNode) *ListNode {
	return merge_20220905(lists, 0, len(lists)-1)
}

func merge_20220905(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	if left > right {
		return nil
	}

	mid := (left + right) / 2
	return mergeTwoList_20220905(merge_20220905(lists, left, mid), merge_20220905(lists, mid+1, right))
}

func mergeTwoList_20220905(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	tail := head
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

	return head.Next
}
