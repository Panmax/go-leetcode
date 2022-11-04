package main

func getBucketId(x, size int) int {
	if x >= 0 {
		return x / size
	}
	return (x+1)/size - 1
}

func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	bucketSize := valueDiff + 1
	m := make(map[int]int)
	for i, num := range nums {
		loc := getBucketId(num, bucketSize)
		if _, exist := m[loc]; exist {
			return true
		}
		// <=
		if x, exist := m[loc-1]; exist && abs(x-num) <= valueDiff {
			return true
		}
		if x, exist := m[loc+1]; exist && abs(x-num) <= valueDiff {
			return true
		}
		m[loc] = num
		if i >= indexDiff {
			// nums[i-indexDiff]
			delete(m, getBucketId(nums[i-indexDiff], bucketSize))
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
