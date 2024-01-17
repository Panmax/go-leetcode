package main

func jump_20240117(nums []int) int {
	var step, end, maxPos int
	for i, num := range nums {
		if i == len(nums)-1 {
			break
		}
		maxPos = max(maxPos, i+num)
		if i == end {
			end = maxPos
			step++
		}
	}
	return step
}
