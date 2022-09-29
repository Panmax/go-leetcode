package main

import "math"

func maximumGap_20220929(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min := math.MaxInt32
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	bucketSize := Max(1, (max-min)/(len(nums)-1))
	buckets := make([]*Bucket, (max-min)/bucketSize+1)
	for _, num := range nums {
		loc := (num - min) / bucketSize
		if buckets[loc] == nil {
			buckets[loc] = NewBucket()
		}
		if num > buckets[loc].Max {
			buckets[loc].Max = num
		}
		if num < buckets[loc].Min {
			buckets[loc].Min = num
		}
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
