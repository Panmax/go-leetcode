package main

func reorderList_20231212(head *ListNode) {
	var list []*ListNode
	node := head
	for node != nil {
		list = append(list, node)
		node = node.Next
	}
	i, j := 0, len(list)-1
	for i < j {
		list[i].Next = list[j]
		i++
		if i >= j {
			break
		}
		list[j].Next = list[i]
		j--
	}
	list[i].Next = nil
}
