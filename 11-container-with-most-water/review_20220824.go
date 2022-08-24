package main

func maxArea_20220824(height []int) int {
	left, right := 0, len(height)-1
	mxArea := 0
	for left < right {
		var area int
		if height[left] < height[right] {
			area = (right - left) * height[left]
			left++
		} else {
			area = (right - left) * height[right]
			right--
		}
		if area > mxArea {
			mxArea = area
		}
	}
	return mxArea
}
