package main

func preorder_20220901(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		// 判断退出条件
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		if len(node.Children) > 0 {
			for _, child := range node.Children {
				dfs(child)
			}
		}
	}
	dfs(root)
	return ans
}
