package main

import "math"

func maximumGap_20220925(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var min, max int
	min = math.MaxInt32
	max = math.MinInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
