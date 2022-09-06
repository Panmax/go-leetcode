package main

func maxArea_20220906(height []int) int {
	mArea := 0
	i, j := 0, len(height)-1
	for i < j {
		area := 0
		if height[i] < height[j] {
			area = (j - i) * height[i]
			i++
		} else {
			area = (j - i) * height[j]
			j--
		}
		if area > mArea {
			mArea = area
		}
	}
	return mArea
}
