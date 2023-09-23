package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	var dfs func(node *TreeNode, preSum int) int
	dfs = func(node *TreeNode, preSum int) int {
		if node == nil {
			// 0
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
