package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
