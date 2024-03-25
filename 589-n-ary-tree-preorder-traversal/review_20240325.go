package main

func preorder_20230325(root *Node) []int {
	var res []int
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
