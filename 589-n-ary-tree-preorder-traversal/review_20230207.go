package main

func preorder_20230207(root *Node) []int {
	res := make([]int, 0)
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return res
}
