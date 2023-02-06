package main

func maxDepth_20230203(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth_20230203(root.Left), maxDepth_20230203(root.Right)) + 1
}
