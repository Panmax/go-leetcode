package main

func containsNearbyAlmostDuplicate_20221103(nums []int, indexDiff int, valueDiff int) bool {
	bucketSize := valueDiff + 1
	m := make(map[int]int)
	for i, num := range nums {
		loc := getBucketId(num, bucketSize)
		if _, exist := m[loc]; exist {
			return true
		}
		if x, exist := m[loc-1]; exist && abs(x-num) <= valueDiff {
			return true
		}
		if x, exist := m[loc+1]; exist && abs(x-num) <= valueDiff {
			return true
		}
		m[loc] = num
		if i >= indexDiff {
			delete(m, getBucketId(nums[i-indexDiff], bucketSize))
		}
	}
	return false
}
