package main

func invertTree_20240106(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	dfs(root)
	return root
}
