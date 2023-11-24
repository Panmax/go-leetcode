package main

func isSymmetric_20231124(root *TreeNode) bool {
	return check_20231124(root.Left, root.Right)
}

func check_20231124(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check_20231124(left.Left, right.Right) && check_20231124(left.Right, right.Left)
}
