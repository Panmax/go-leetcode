package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	ans := 1
	var depth func(*TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := depth(node.Left)
		r := depth(node.Right)
		ans = max(ans, r+l+1)
		return max(r, l) + 1
	}
	depth(root)
	return ans - 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
