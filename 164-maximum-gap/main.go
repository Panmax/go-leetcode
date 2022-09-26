package main

import (
	"math"
)

type Bucket struct {
	Min int
	Max int
}

func NewBucket() *Bucket {
	return &Bucket{
		Min: math.MaxInt32,
		Max: math.MinInt32,
	}
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	min := math.MaxInt32
	max := math.MinInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	bucketSize := (max - min) / (len(nums) - 1)
	if bucketSize < 1 {
		bucketSize = 1
	}
	buckets := make([]*Bucket, (max-min)/bucketSize+1)
	for _, num := range nums {
		loc := (num - min) / bucketSize
		if buckets[loc] == nil {
			buckets[loc] = NewBucket()
		}
		if buckets[loc].Min > num {
			buckets[loc].Min = num
		}
		if buckets[loc].Max < num {
			buckets[loc].Max = num
		}
	}

	previousMax := math.MaxInt32
	maxGap := math.MinInt32
	for _, bucket := range buckets {
		if bucket != nil && previousMax != math.MaxInt32 {
			if maxGap < bucket.Min-previousMax {
				maxGap = bucket.Min - previousMax
			}
		}
		if bucket != nil {
			previousMax = bucket.Max
			if maxGap < bucket.Max-bucket.Min {
				maxGap = bucket.Max - bucket.Min
			}
		}
	}
	return maxGap
}
