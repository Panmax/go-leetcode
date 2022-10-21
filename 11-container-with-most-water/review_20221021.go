package main

func maxArea_20221021(height []int) int {
	mArea := 0
	left, right := 0, len(height)-1
	for left < right {
		area := 0
		if height[left] < height[right] {
			area = height[left] * (right - left)
			left++
		} else {
			area = height[right] * (right - left)
			right--
		}
		if area > mArea {
			mArea = area
		}
	}
	return mArea
}
