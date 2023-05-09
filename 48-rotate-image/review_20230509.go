package main

func rotate_20230509(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		// j := i
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
