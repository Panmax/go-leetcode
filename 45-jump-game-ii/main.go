package main

func jump(nums []int) int {
	var step, end, maxPos int
	for i, num := range nums {
		// 不需要计算最后一个位置
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

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
