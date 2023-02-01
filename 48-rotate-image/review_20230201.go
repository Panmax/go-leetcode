package main

func rotate_20230201(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		// j := i
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[n-1-j][i], matrix[n-i-1][n-j-1], matrix[j][n-1-i] =
				matrix[n-1-j][i], matrix[n-i-1][n-j-1], matrix[j][n-1-i], matrix[i][j]
		}
	}
}
