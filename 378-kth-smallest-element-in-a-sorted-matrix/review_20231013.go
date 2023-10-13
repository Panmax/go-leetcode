package main

func kthSmallest_20231013(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		if isInLeft_20231013(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func isInLeft_20231013(matrix [][]int, mid int, k int, n int) bool {
	var count int
	i, j := n-1, 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			// += i + 1
			count += i + 1
			j++
		} else {
			i--
		}
	}
	return count >= k
}
