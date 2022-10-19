package main

import "math"

func maxPathSum_20221019(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := max(0, maxGain(node.Left))
		rightMax := max(0, maxGain(node.Right))
		maxSum = max(maxSum, node.Val+leftMax+rightMax)
		return node.Val + max(leftMax, rightMax)
	}
	maxGain(root)
	return maxSum
}
