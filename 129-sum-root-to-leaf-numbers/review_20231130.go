package main

func sumNumbers_20231130(root *TreeNode) int {
	var dfs func(node *TreeNode, preSum int) int
	dfs = func(node *TreeNode, preSum int) int {
		if node == nil {
			return 0
		}
		preSum = preSum*10 + node.Val
		if node.Left == nil && node.Right == nil {
			return preSum
		}

		return dfs(node.Left, preSum) + dfs(node.Right, preSum)
	}
	return dfs(root, 0)
}
