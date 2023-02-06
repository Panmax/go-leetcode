package main

func recoverTree_20230203(root *TreeNode) {
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
	x, y := findTwoSwitch_20230203(nums)
	recovery_20230203(root, 2, x, y)
}

func findTwoSwitch_20230203(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index2 = i + 1
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	// 这里返回的是数字，不是index
	return nums[index1], nums[index2]
}

func recovery_20230203(root *TreeNode, n, x, y int) {
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
	recovery_20230203(root.Left, n, x, y)
	recovery_20230203(root.Right, n, x, y)
}
