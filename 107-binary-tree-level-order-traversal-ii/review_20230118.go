package main

func levelOrderBottom_20230118(root *TreeNode) [][]int {
	var ans [][]int
	var queue []*TreeNode
	if root == nil {
		return ans
	}
	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		n := len(queue)
		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			level = append(level, node.Val)
		}
		ans = append(ans, level)
	}

	for i := 0; i < len(ans)/2; i++ {
		ans[i], ans[len(ans)-i-1] = ans[len(ans)-i-1], ans[i]
	}
	return ans
}
