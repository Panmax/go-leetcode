package main

func zigzagLevelOrder_20230821(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode
	queue = append(queue, root)

	var level int
	for len(queue) > 0 {
		level++
		n := len(queue)
		var vals []int

		for i := 0; i < n; i++ {
			node := queue[0]
			queue = queue[1:]
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 0 {
			for i := 0; i < len(vals)/2; i++ {
				vals[i], vals[len(vals)-i-1] = vals[len(vals)-i-1], vals[i]
			}
		}
		res = append(res, vals)
	}
	return res
}
