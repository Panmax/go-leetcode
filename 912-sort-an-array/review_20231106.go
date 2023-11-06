package main

func sortArray_20231106(nums []int) []int {
	quickSort_20231106(nums)
	return nums
}

func quickSort_20231106(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	pivot := nums[0]
	i, j := -1, n
	for i < j {
		for i++; nums[i] < pivot; i++ {
		}
		for j--; nums[j] > pivot; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	quickSort_20231106(nums[:j+1])
	quickSort_20231106(nums[j+1:])
}
