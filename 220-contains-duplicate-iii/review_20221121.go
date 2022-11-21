package main

func containsNearbyAlmostDuplicate_20221121(nums []int, indexDiff int, valueDiff int) bool {
	size := valueDiff + 1
	m := make(map[int]int)
	for i, num := range nums {
		loc := getBucketId(num, size)
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
		// >=
		if i >= indexDiff {
			delete(m, getBucketId(nums[i-indexDiff], size))
		}
	}
	return false
}
