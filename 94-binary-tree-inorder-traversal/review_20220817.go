package main

func inorderTraversal_20220817(root *TreeNode) []int {
	var ans []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		// 这里一定要注意，每次都忘判断
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return ans
}
