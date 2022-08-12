package main

func twoSum_20220812(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if j, ok := numMap[target-num]; ok {
			return []int{i, j}
		}
		numMap[num] = i
	}
	return nil
}
