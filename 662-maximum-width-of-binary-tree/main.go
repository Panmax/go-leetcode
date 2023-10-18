package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type pair struct {
	node  *TreeNode
	index int
}

func widthOfBinaryTree(root *TreeNode) int {
	ans := 1
	var queue []pair
	queue = append(queue, pair{
		node:  root,
		index: 1,
	})
	for len(queue) > 0 {
		length := len(queue)
		ans = max(ans, queue[length-1].index-queue[0].index+1)
		for i := 0; i < length; i++ {
			if queue[i].node.Left != nil {
				queue = append(queue, pair{
					node:  queue[i].node.Left,
					index: queue[i].index * 2,
				})
			}
			if queue[i].node.Right != nil {
				queue = append(queue, pair{
					node:  queue[i].node.Right,
					index: queue[i].index*2 + 1,
				})
			}
		}

		queue = queue[length:]
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
