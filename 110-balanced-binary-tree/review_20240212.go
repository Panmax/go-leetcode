package main

func isBalanced_20240212(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height_20240212(root.Left)-height_20240212(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height_20240212(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height_20240212(node.Left), height_20240212(node.Right)) + 1
}
