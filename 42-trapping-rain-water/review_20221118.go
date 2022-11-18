package main

func trap_20221118(height []int) int {
	var stack []int
	res := 0
	for i, rightHeight := range height {
		for len(stack) > 0 && rightHeight > height[stack[len(stack)-1]] {
			bottomHeight := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			leftIndex := stack[len(stack)-1]
			leftHeight := height[leftIndex]
			res += (i - leftIndex - 1) * (min(leftHeight, rightHeight) - bottomHeight)
		}
		stack = append(stack, i)
	}
	return res
}
