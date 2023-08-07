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
	q := randomPartition(nums, l, r)
	if q == index {
		return nums[q]
	} else if q < index {
		return quickSelect(nums, q+1, r, index)
	} else {
		return quickSelect(nums, l, q-1, index)
	}
}

func randomPartition(nums []int, l int, r int) int {
	// r-l+1
	i := rand.Intn(r-l+1) + l
	nums[i], nums[r] = nums[r], nums[i]
	x := nums[r]
	i = l - 1
	// j<r
	for j := l; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
