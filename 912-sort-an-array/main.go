package main

func sortArray(nums []int) []int {
	quickSort(nums)
	return nums
}

func quickSort(nums []int) {
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
	quickSort(nums[:j+1])
	quickSort(nums[j+1:])
}
