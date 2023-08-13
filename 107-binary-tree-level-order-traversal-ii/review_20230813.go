package main

func levelOrderBottom_20230813(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int

	var queue []*TreeNode
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

		res = append(res, level)
	}
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}
