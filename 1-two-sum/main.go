package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if index, ok := numMap[target-num]; ok {
			return []int{i, index}
		}
		numMap[num] = i
	}
	return nil
}
