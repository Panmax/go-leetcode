package main

import (
	"math"
	"sort"
)

func getStrongest(arr []int, k int) []int {
	strongest := make([]int, k)
	n := len(arr)
	sort.Ints(arr)
	m := arr[(n-1)/2]
	left, right := 0, n-1
	for i := 0; i < k; i++ {
		if leftStrong(m, arr[left], arr[right]) {
			strongest[i] = arr[left]
			left++
		} else {
			strongest[i] = arr[right]
			right--
		}
	}
	return strongest
}

func leftStrong(m, leftNum, rightNum int) bool {
	if math.Abs(float64(leftNum-m)) > math.Abs(float64(rightNum-m)) {
		return true
	} else { // 其余情况都是 false，因为已经排序，相等时一定是右边的数更大
		return false
	}
}
