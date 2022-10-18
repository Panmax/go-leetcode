package main

func trap_20221018(height []int) int {
	var stack []int
	res := 0
	for i, rightHeight := range height {
		// for
		for len(stack) > 0 && rightHeight > height[stack[len(stack)-1]] {
			bottomIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIndex := stack[len(stack)-1]
			leftHeight := height[leftIndex]
			bottomHeight := height[bottomIndex]
			res += (i - leftIndex - 1) * (min(leftHeight, rightHeight) - bottomHeight)
		}
		stack = append(stack, i)
	}
	return res
}
