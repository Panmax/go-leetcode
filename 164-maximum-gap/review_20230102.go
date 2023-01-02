package main

import "math"

func maximumGap_20230102(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	max := math.MinInt32
	min := math.MaxInt32
	for _, num := range nums {
		max = Max(max, num)
		min = Min(min, num)
	}
	bucketSize := Max(1, (max-min)/(len(nums)-1))
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
	maxGap := 0
	for _, bucket := range buckets {
		if bucket != nil && preMax != math.MaxInt32 {
			maxGap = Max(maxGap, bucket.Min-preMax)
		}
		if bucket != nil {
			maxGap = Max(maxGap, bucket.Max-bucket.Min)
			// bucket.Max
			preMax = bucket.Max
		}
	}
	return maxGap
}
