package main

import (
	"math"
	"sort"
)

func getStrongest_20231117(arr []int, k int) []int {
	strongest := make([]int, k)
	n := len(arr)
	// sort
	sort.Ints(arr)
	m := arr[(n-1)/2]
	left, right := 0, n-1
	for i := 0; i < k; i++ {
		if isLeftStrong_20231117(m, arr[left], arr[right]) {
			strongest[i] = arr[left]
			left++
		} else {
			strongest[i] = arr[right]
			right--
		}
	}

	return strongest
}

func isLeftStrong_20231117(m int, left int, right int) bool {
	return math.Abs(float64(left-m)) > math.Abs(float64(right-m))
}
