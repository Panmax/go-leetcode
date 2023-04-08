package main

func flatten_20230409(root *TreeNode) {
	list := make([]*TreeNode, 0)
	var preorder func(node *TreeNode)
	preorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		list = append(list, node)
		preorder(node.Left)
		preorder(node.Right)
	}
	preorder(root)
	for i := 1; i < len(list); i++ {
		pre, cur := list[i-1], list[i]
		pre.Left, pre.Right = nil, cur
	}
}
