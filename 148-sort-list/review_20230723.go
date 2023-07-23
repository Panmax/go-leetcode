package main

func sortList_20230723(head *ListNode) *ListNode {
	return sort_20230723(head, nil)
}

func sort_20230723(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	mid := head
	fast := head
	for fast != tail && fast.Next != tail {
		mid = mid.Next
		fast = fast.Next.Next
	}
	return merge_20230723(sort_20230723(head, mid), sort_20230723(mid, tail))
}

func merge_20230723(node1 *ListNode, node2 *ListNode) *ListNode {
	dummy := &ListNode{}
	temp, temp1, temp2 := dummy, node1, node2
	for temp1 != nil && temp2 != nil {
		if temp1.Val < temp2.Val {
			temp.Next = temp1
			temp1 = temp1.Next
		} else {
			temp.Next = temp2
			temp2 = temp2.Next
		}
		temp = temp.Next
	}
	if temp1 != nil {
		temp.Next = temp1
	} else {
		temp.Next = temp2
	}
	return dummy.Next
}
