package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	rand.Seed(time.Now().UnixNano())
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, l int, r int, index int) int {
	q := randomPartition(nums, l, r)
	if q == index {
		return nums[q]
	} else if q < index { // 在右边找
		return quickSelect(nums, q+1, r, index)
	}
	// 在左边找
	return quickSelect(nums, l, q-1, index)
}

func randomPartition(nums []int, l int, r int) int {
	// 随机找个锚点，需要+l
	i := rand.Intn(r-l+1) + l
	// 把锚点值临时放到最右边
	nums[i], nums[r] = nums[r], nums[i]
	x := nums[r]

	// 把小于锚点之的放在左边
	i = l - 1
	for j := l; j < r; j++ {
		if nums[j] <= x {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	// 锚点之放回原位
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}
