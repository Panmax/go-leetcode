package main

func inorderTraversal_20221016(root *TreeNode) []int {
	var inorder func(node *TreeNode)
	var res []int
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 不需要判断 if
		// if node.Left != nil {
		inorder(node.Left)
		// }
		res = append(res, node.Val)
		// if node.Right != nil {
		inorder(node.Right)
		// }
	}
	inorder(root)
	return res
}
