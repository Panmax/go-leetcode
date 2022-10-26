package main

func trap_20221026(height []int) int {
	res := 0
	var stack []int
	for i, rightHeight := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < rightHeight {
			bottomIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// if
			if len(stack) == 0 {
				break
			}
			bottomHeight := height[bottomIndex]
			leftIndex := stack[len(stack)-1]
			leftHeight := height[leftIndex]
			// min(leftHeight, rightHeight) - bottomHeight
			res += (i - leftIndex - 1) * (min(leftHeight, rightHeight) - bottomHeight)
		}
		stack = append(stack, i)
	}
	return res
}
