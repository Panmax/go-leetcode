package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parents := make(map[*TreeNode]*TreeNode)
	visited := make(map[*TreeNode]bool)

	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parents[node.Left] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parents[node.Right] = node
			dfs(node.Right)
		}
	}
	dfs(root)

	for p != nil {
		visited[p] = true
		p = parents[p]
	}
	for q != nil {
		if visited[q] {
			return q
		}
		q = parents[q]
	}
	return nil
}
