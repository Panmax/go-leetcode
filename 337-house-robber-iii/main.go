package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	m1 := make(map[*TreeNode]int) // 存储选了当前节点，不选子节点的结果
	m2 := make(map[*TreeNode]int) // 存储不选当前节点，选子节点的结果
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		m1[node] = node.Val + m2[node.Left] + m2[node.Right]
		m2[node] = Max(m1[node.Left], m2[node.Left]) + Max(m1[node.Right], m2[node.Right])
	}
	dfs(root)
	return Max(m1[root], m2[root])
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
