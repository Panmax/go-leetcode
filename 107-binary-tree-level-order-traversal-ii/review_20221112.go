package main

func levelOrderBottom_20221112(root *TreeNode) [][]int {
	var res [][]int
	var queue []*TreeNode
	if root == nil {
		return res
	}

	queue = append(queue, root)
	for len(queue) > 0 {
		var level []int
		length := len(queue)
		for i := 0; i < length; i++ {
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
