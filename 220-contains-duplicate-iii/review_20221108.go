package main

func containsNearbyAlmostDuplicate_20221108(nums []int, indexDiff int, valueDiff int) bool {
	bucketSize := valueDiff + 1
	m := make(map[int]int)
	for i, num := range nums {
		loc := getBucketId(num, bucketSize)
		if _, ok := m[loc]; ok {
			return true
		}
		if n, ok := m[loc-1]; ok && abs(num-n) <= valueDiff {
			return true
		}
		if n, ok := m[loc+1]; ok && abs(num-n) <= valueDiff {
			return true
		}
		m[loc] = num
		if i >= indexDiff {
			delete(m, getBucketId(nums[i-indexDiff], bucketSize))
		}
	}
	return false
}
