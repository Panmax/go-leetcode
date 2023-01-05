package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var nums []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		nums = append(nums, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	x, y := findTwoSwiped(nums)
	recover(root, 2, x, y)
}

func findTwoSwiped(nums []int) (int, int) {
	index1, index2 := -1, -1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			index2 = i + 1
			// if 在里边
			if index1 == -1 {
				index1 = i
			} else {
				break
			}
		}
	}
	return nums[index1], nums[index2]
}

func recover(node *TreeNode, count, x, y int) {
	if node == nil {
		return
	}
	if node.Val == x || node.Val == y {
		if node.Val == x {
			node.Val = y
		} else { // else
			node.Val = x
		}
		count--
		if count == 0 {
			return
		}
	}
	recover(node.Left, count, x, y)
	recover(node.Right, count, x, y)
}
