package main

func preorder_20220811(root *Node) (ans []int) {
	var helpFunc func(*Node)
	helpFunc = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for i := range node.Children {
			helpFunc(node.Children[i])
		}
	}

	helpFunc(root)
	return
}
