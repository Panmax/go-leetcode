package main

func candy(ratings []int) int {
	n := len(ratings)
	var ans int

	left := make([]int, n)
	for i := range ratings {
		if i > 0 && ratings[i] > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}

	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
