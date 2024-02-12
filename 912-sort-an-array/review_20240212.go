package main

func sortArray_20240212(nums []int) []int {
	quickSort_20240212(nums)
	return nums
}

func quickSort_20240212(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j := -1, n
	pivot := nums[0]
	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort_20240212(nums[:j+1])
	quickSort_20240212(nums[j+1:])
}
