package main

func isPalindrome_20230317(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}
	reverseX := 0
	for x > reverseX {
		reverseX = reverseX*10 + x%10
		x /= 10
	}
	return x == reverseX || x == reverseX/10
}
