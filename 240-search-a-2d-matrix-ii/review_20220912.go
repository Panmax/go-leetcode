package main

func searchMatrix_20220912(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i, j := 0, n-1
	for i < m && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}
