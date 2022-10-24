package main

import "math"

func maximumGap_20221024(nums []int) int {
	min := math.MaxInt32
	max := math.MinInt32
	for _, num := range nums {
		max = Max(num, max)
		min = Min(num, min)
	}

	bucketSize := Max(1, (max-min)/len(nums))
	buckets := make([]*Bucket, (max-min)/bucketSize+1)
	for _, num := range nums {
		loc := (num - min) / bucketSize
		if buckets[loc] == nil {
			buckets[loc] = NewBucket()
		}
		buckets[loc].Max = Max(buckets[loc].Max, num)
		buckets[loc].Min = Min(buckets[loc].Min, num)
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
