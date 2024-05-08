package main

func isSymmetric_20240508(root *TreeNode) bool {
	return check_20240508(root.Left, root.Right)
}

func check_20240508(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && check_20240508(left.Left, right.Right) && check_20240508(left.Right, right.Left)
}
