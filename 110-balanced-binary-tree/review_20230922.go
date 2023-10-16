package main

func isBalanced_20230922(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height_20230922(root.Left)-height_20230922(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height_20230922(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return max(height_20230922(node.Left), height_20230922(node.Right)) + 1
}
