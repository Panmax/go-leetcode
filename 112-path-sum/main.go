package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	// 叶子节点
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	targetSum = targetSum - root.Val
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
