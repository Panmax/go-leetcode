package main

func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	x, y := 0, n-1
	for x < m && y >= 0 {
		if matrix[x][y] == target {
			return true
		} else if matrix[x][y] > target {
			y--
		} else {
			x++
		}
	}
	return false
}

func searchMatrixLoop(matrix [][]int, target int) bool {
	for i := range matrix {
		if binarySearch(matrix[i], target) {
			return true
		}
	}
	return false
}

func binarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right + left) / 2
		if nums[mid] == target {
			return true
		}
		if nums[left] <= target && nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
