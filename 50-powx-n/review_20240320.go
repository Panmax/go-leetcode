package main

func myPow_20240320(x float64, n int) float64 {
	if n > 0 {
		return quickMul_20240320(x, n)
	}
	return 1 / quickMul_20240320(x, n)
}

func quickMul_20240320(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMul_20240320(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
