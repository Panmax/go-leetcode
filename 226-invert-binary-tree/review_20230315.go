package main

func invertTree_20230315(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 判断 nil
		dfs(node.Left)
		dfs(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	dfs(root)
	return root
}
