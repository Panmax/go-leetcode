package main

func rotate_20230131(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ { //{i , j}-->{j , n-1-i}-->{n-1-i , n-1-j}-->{n-1-j ,i}--> {i , j}
		for j := i; j < n-1-i; j++ {
			matrix[i][j], matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i] =
				matrix[n-1-j][i], matrix[n-1-i][n-1-j], matrix[j][n-1-i], matrix[i][j]
		}
	}
}
