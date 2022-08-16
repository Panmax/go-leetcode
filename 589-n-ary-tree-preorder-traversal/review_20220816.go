package main

func preorder_20220816(root *Node) (ans []int) {
	var pre func(node *Node)
	pre = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			pre(child)
		}
	}
	pre(root)
	return
}
