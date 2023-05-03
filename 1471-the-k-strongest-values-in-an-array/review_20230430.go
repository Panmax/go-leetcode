package main

import (
	"math"
	"sort"
)

func getStrongest_20230429(arr []int, k int) []int {
	res := make([]int, 0, k)
	n := len(arr)
	sort.Ints(arr)
	m := arr[(n-1)/2]
	left, right := 0, n-1
	for len(res) < k {
		// arr[left], arr[right]
		if isLeftStrong_20230429(m, arr[left], arr[right]) {
			res = append(res, arr[left])
			left++
		} else {
			res = append(res, arr[right])
			right--
		}
	}

	return res
}

func isLeftStrong_20230429(mid, left, right int) bool {
	if math.Abs(float64(left-mid)) > math.Abs(float64(right-mid)) {
		return true
	}
	return false
}
