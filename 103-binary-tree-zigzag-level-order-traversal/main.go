package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var res [][]int

	queue := []*TreeNode{root}
	level := 0
	for len(queue) > 0 {
		q := queue
		queue = nil
		var vals []int
		for _, node := range q {
			vals = append(vals, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if level%2 == 1 {
			// len(vals)/2
			for i := 0; i < len(vals)/2; i++ {
				vals[i], vals[len(vals)-i-1] = vals[len(vals)-i-1], vals[i]
			}
		}
		res = append(res, vals)
		level++
	}

	return res
}
