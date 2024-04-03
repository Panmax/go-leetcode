package main

func diameterOfBinaryTree_20240403(root *TreeNode) int {
	ans := 1
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		ans = max(ans, l+r+1)
		return max(r, l) + 1
	}
	depth(root)
	return ans - 1
}
