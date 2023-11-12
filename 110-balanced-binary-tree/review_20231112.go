package main

func isBalanced_20231112(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height_20231112(root.Left)-height_20231112(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height_20231112(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height_20231112(node.Left), height_20231112(node.Right)) + 1
}
