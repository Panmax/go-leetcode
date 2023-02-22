package main

import (
	"math"
	"sort"
)

func getStrongest_20230222(arr []int, k int) []int {
	res := make([]int, k)
	n := len(arr)
	sort.Ints(arr)
	m := arr[(n-1)/2]
	left, right := 0, n-1
	for i := 0; i < k; i++ {
		if isLeftStrong_20230222(m, arr[left], arr[right]) {
			res[i] = arr[left]
			left++
		} else {
			res[i] = arr[right]
			right--
		}
	}
	return res
}

func isLeftStrong_20230222(m int, left int, right int) bool {
	if math.Abs(float64(left-m)) > math.Abs(float64(right-m)) {
		return true
	}
	return false
}
