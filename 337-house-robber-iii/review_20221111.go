package main

func rob_20221111(root *TreeNode) int {
	includeMap := make(map[*TreeNode]int)
	excludeMap := make(map[*TreeNode]int)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		includeMap[node] = node.Val + excludeMap[node.Left] + excludeMap[node.Right]
		excludeMap[node] = Max(includeMap[node.Left], excludeMap[node.Left]) + Max(includeMap[node.Right], excludeMap[node.Right])
	}

	dfs(root)
	return Max(includeMap[root], excludeMap[root])
}
