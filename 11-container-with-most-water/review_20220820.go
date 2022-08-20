package main

func maxArea_20220820(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		area := 0
		if height[left] > height[right] {
			area = (right - left) * height[right]
			right--
		} else {
			area = (right - left) * height[left]
			left++
		}
		if area > res {
			res = area
		}
	}

	return res
}
