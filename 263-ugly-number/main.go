package main

var factors = []int{2, 3, 5}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, factor := range factors {
		for n%factor == 0 {
			n /= factor
		}
	}
	return n == 1
}
