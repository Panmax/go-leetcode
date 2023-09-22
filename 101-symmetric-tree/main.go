package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return check(root, root)
}

func check(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	// 比较值 && (左的左 和 右的右) && (左的右 和 右的左)
	return l.Val == r.Val && check(l.Left, r.Right) && check(l.Right, r.Left)
}
