package main

func reorderList_20230821(head *ListNode) {
	var nodes []*ListNode
	node := head
	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i >= j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	// 最后切断
	nodes[i].Next = nil
}
