package main

func kthSmallest_20230129(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if isInLeft_20230129(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isInLeft_20230129(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	var num int
	for i >= 0 && j < n {
		// <=
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	return num >= k
}
