package main

func maxDepth_20230123(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth_20230123(root.Left), maxDepth_20230123(root.Right)) + 1
}
