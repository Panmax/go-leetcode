package main

func rob_20221212(root *TreeNode) int {
	includeMap := make(map[*TreeNode]int)
	excludeMap := make(map[*TreeNode]int)
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		// node.Val +
		includeMap[node] = node.Val + excludeMap[node.Left] + excludeMap[node.Right]
		// Left,Left  Right,Right
		excludeMap[node] = Max(excludeMap[node.Left], includeMap[node.Left]) + Max(excludeMap[node.Right], includeMap[node.Right])
	}
	dfs(root)
	return Max(includeMap[root], excludeMap[root])
}
