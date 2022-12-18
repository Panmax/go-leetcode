package main

func increasingTriplet_20221218(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}
	leftMin := make([]int, n)
	leftMin[0] = nums[0]
	for i := 1; i < n-1; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
	}

	rightMax := make([]int, n)
	rightMax[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	for i := 1; i <= n-2; i++ {
		if leftMin[i] < nums[i] && rightMax[i] > nums[i] {
			return true
		}
	}
	return false
}
