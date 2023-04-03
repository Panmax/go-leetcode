package main

func recoverTree_20230403(root *TreeNode) {
	var nums []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwitch_20230403(nums)
	recovery_20230403(root, 2, x, y)
}

func findTwoSwitch_20230403(nums []int) (x, y int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			}
		}
	}
	return nums[index1], nums[index2]
}

func recovery_20230403(root *TreeNode, n int, x int, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else if root.Val == y {
			root.Val = x
		}
		n--
		if n <= 0 {
			return
		}
	}
	recovery_20230403(root.Left, n, x, y)
	recovery_20230403(root.Right, n, x, y)
}
