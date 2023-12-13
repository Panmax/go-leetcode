package main

func isUgly_20231213(n int) bool {
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
