package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var list []*TreeNode
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
