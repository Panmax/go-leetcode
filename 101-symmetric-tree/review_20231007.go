package main

func isSymmetric_20231007(root *TreeNode) bool {
	return check_20231007(root, root)
}

func check_20231007(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && check_20231007(l.Left, r.Right) && check_20231007(l.Right, r.Left)
}
