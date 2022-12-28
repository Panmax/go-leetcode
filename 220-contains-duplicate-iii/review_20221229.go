package main

func getBucket(num, size int) int {
	if num >= 0 {
		return num / size
	}
	return (num+1)/size - 1
}

func containsNearbyAlmostDuplicate_20221229(nums []int, indexDiff int, valueDiff int) bool {
	// valueDiff
	size := valueDiff + 1
	m := make(map[int]int)
	for i, num := range nums {
		loc := getBucket(num, size)
		if _, ok := m[loc]; ok {
			return true
		}
		if x, ok := m[loc-1]; ok {
			if abs(x-num) <= valueDiff {
				return true
			}
		}
		if x, ok := m[loc+1]; ok {
			if abs(x-num) <= valueDiff {
				return true
			}
		}
		m[loc] = num
		if i >= indexDiff {
			delete(m, getBucket(nums[i-indexDiff], size))
		}
	}
	return false
}
