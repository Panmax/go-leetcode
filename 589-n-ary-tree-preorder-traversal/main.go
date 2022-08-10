package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	result := make([]int, 0)
	for {
		result = append(result, root.Val)
		for i := range root.Children {
			result = append(result, preorder(root.Children[i])...)
		}
		break
	}

	return result
}

func preorder_Official(root *Node) (ans []int) {
	var dfs func(*Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return ans
}
