package main

func containsNearbyAlmostDuplicate_20230518(nums []int, indexDiff int, valueDiff int) bool {
	size := valueDiff + 1
	m := make(map[int]int)
	for i, num := range nums {
		loc := getBucket(num, size)
		if _, ok := m[loc]; ok {
			return true
		}
		if _, ok := m[loc-1]; ok && abs(m[loc-1]-num) <= valueDiff {
			return true
		}
		if _, ok := m[loc+1]; ok && abs(m[loc+1]-num) <= valueDiff {
			return true
		}
		m[loc] = num

		if i >= indexDiff {
			delete(m, getBucket(nums[i-indexDiff], size))
		}
	}
	return false
}
