package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l, r int, index int) int {
	i := randomPartition(nums, l, r)
	if i == index {
		return nums[i]
	} else if i < index {
		return quickSelect(nums, i+1, r, index)
	} else {
		return quickSelect(nums, l, i-1, index)
	}
}

func randomPartition(nums []int, l int, r int) int {
	i := rand.Intn(r-l+1) + l
	x := nums[i]
	nums[i], nums[r] = nums[r], nums[i]
	i = l - 1
	for j := l; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
