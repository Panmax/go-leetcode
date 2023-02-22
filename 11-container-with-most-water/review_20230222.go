package main

func maxArea_20230222(height []int) int {
	var max int
	left, right := 0, len(height)-1
	for left < right {
		var area int
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if area > max {
			max = area
		}
	}

	return max
}
