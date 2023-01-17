package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var max int

	var findMaxLevel func(node *TreeNode, level int)
	findMaxLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		level++
		if level > max {
			max = level
		}
		findMaxLevel(node.Left, level)
		findMaxLevel(node.Right, level)
	}

	findMaxLevel(root, 0)
	return max
}
