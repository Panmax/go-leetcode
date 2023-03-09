package main

import (
	"math"
	"sort"
)

func getStrongest_20230309(arr []int, k int) []int {
	res := make([]int, k)
	sort.Ints(arr)
	n := len(arr)
	mid := arr[(n-1)/2]
	left, right := 0, n-1
	for i := 0; i < k; i++ {
		if isLeftStrong_20230309(mid, arr[left], arr[right]) {
			res[i] = arr[left]
			left++
		} else {
			res[i] = arr[right]
			right--
		}
	}
	return res
}

func isLeftStrong_20230309(mid int, left int, right int) bool {
	if math.Abs(float64(left-mid)) > math.Abs(float64(right-mid)) {
		return true
	}
	return false
}
