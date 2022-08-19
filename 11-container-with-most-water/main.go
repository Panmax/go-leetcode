package main

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		width := right - left
		var area int
		if height[left] <= height[right] {
			area = width * height[left]
			left++
		} else {
			area = width * height[right]
			right--
		}
		if area > res {
			res = area
		}
	}
	return res
}
