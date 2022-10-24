package main

import "math"

func maxPathSum_20221024(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := max(maxGain(node.Left), 0)
		rightMax := max(maxGain(node.Right), 0)
		maxSum = max(maxSum, node.Val+leftMax+rightMax)
		return max(node.Val+leftMax, node.Val+rightMax)
	}
	maxGain(root)
	return maxSum
}
