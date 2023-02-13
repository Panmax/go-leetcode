package main

func trap_20230212(height []int) int {
	var stack []int
	var res int
	for i, rightHeight := range height {
		for len(stack) > 0 && rightHeight > height[stack[len(stack)-1]] {
			bottomIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			bottomHeight := height[bottomIndex]
			leftIndex := stack[len(stack)-1]
			leftHeight := height[leftIndex]
			res += (i - leftIndex - 1) * (min(leftHeight, rightHeight) - bottomHeight)
		}
		stack = append(stack, i)
	}
	return res
}
