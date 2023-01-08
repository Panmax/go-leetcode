package main

func recoverTree_20230107(root *TreeNode) {
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
	x, y := findTwoSwiped_20230107(nums)
	recovery_20230107(root, 2, x, y)
}

func findTwoSwiped_20230107(nums []int) (int, int) {
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
	return nums[index1], nums[index2]
}

func recovery_20230107(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}
	if root.Val == x || root.Val == y {
		if root.Val == x {
			root.Val = y
		} else {
			root.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recovery_20230107(root.Left, count, x, y)
	recovery_20230107(root.Right, count, x, y)
}
