package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	// math.MinInt32
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := max(maxGain(node.Left), 0)
		rightMax := max(maxGain(node.Right), 0)
		maxSum = max(maxSum, node.Val+leftMax+rightMax)
		return node.Val + max(leftMax, rightMax)
	}
	maxGain(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
