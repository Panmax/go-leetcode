package main

func trap_20221013(height []int) int {
	res := 0
	var stack []int
	for i, rightHeight := range height {
		for len(stack) > 0 && rightHeight > height[stack[len(stack)-1]] {
			bottomIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) <= 0 {
				break
			}
			leftIndex := stack[len(stack)-1]
			leftHeight := height[leftIndex]
			res += (i - leftIndex - 1) * (min(leftHeight, rightHeight) - height[bottomIndex])
		}
		stack = append(stack, i)
	}

	return res
}

func main() {
	print(trap_20221013([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
