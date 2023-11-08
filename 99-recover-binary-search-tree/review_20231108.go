package main

func recoverTree_20231108(root *TreeNode) {
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
	x, y := findTwoSwitch_20231108(nums)
	recovery_20231108(root, 2, x, y)
}

func findTwoSwitch_20231108(nums []int) (x, y int) {
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

func recovery_20231108(node *TreeNode, n int, x int, y int) {
	if node == nil || n == 0 {
		return
	}
	if node.Val == x || node.Val == y {
		if node.Val == x {
			node.Val = y
		} else {
			node.Val = x
		}
		n--
	}
	recovery_20231108(node.Left, n, x, y)
	recovery_20231108(node.Right, n, x, y)
}
