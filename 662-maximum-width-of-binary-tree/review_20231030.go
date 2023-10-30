package main

func widthOfBinaryTree_20231030(root *TreeNode) int {
	ans := 1
	var queue []*pair
	queue = append(queue, &pair{
		node:  root,
		index: 1,
	})
	for len(queue) > 0 {
		n := len(queue)
		// +1
		ans = max(ans, queue[n-1].index-queue[0].index+1)
		for i := 0; i < n; i++ {
			if queue[i].node.Left != nil {
				queue = append(queue, &pair{
					node:  queue[i].node.Left,
					index: queue[i].index * 2,
				})
			}
			if queue[i].node.Right != nil {
				queue = append(queue, &pair{
					node:  queue[i].node.Right,
					index: queue[i].index*2 + 1,
				})
			}
		}
		queue = queue[n:]
	}
	return ans
}
