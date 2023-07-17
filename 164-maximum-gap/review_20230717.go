package main

import "math"

func maximumGap_20230717(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min := math.MaxInt32
	max := math.MinInt32
	for _, num := range nums {
		min = Min(min, num)
		max = Max(max, num)
	}

	bucketSize := (max - min) / (len(nums) - 1)
	bucketSize = Max(1, bucketSize)
	// +1
	buckets := make([]*Bucket, (max-min)/bucketSize+1)
	for _, num := range nums {
		loc := (num - min) / bucketSize
		if buckets[loc] == nil {
			buckets[loc] = NewBucket()
		}
		buckets[loc].Min = Min(buckets[loc].Min, num)
		buckets[loc].Max = Max(buckets[loc].Max, num)
	}
	preMax := math.MaxInt32
	maxGap := math.MinInt32
	for _, bucket := range buckets {
		if bucket != nil && preMax != math.MaxInt32 {
			maxGap = Max(maxGap, bucket.Min-preMax)
		}
		if bucket != nil {
			preMax = bucket.Max
			maxGap = Max(maxGap, bucket.Max-bucket.Min)
		}
	}
	return maxGap
}
