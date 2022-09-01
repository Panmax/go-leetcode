package main

func inorderTraversal_20220901(root *TreeNode) []int {
	ans := make([]int, 0)
	var f func(node *TreeNode)
	f = func(node *TreeNode) {
		if node == nil {
			return
		}
		f(node.Left)
		ans = append(ans, node.Val)
		f(node.Right)
	}
	f(root)
	return ans
}
