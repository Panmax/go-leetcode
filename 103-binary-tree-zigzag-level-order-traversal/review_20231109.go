package main

func zigzagLevelOrder_20231109(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int
	var queue []*TreeNode
	var level int
	queue = append(queue, root)
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
			for i := 0; i < n/2; i++ {
				vals[i], vals[n-i-1] = vals[n-i-1], vals[i]
			}
		}

		res = append(res, vals)

	}
	return res
}
