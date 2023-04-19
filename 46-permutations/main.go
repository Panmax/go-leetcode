package main

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func([]int, []int)
	backtrack = func(nums []int, tmp []int) {
		if len(nums) == 0 {
			res = append(res, tmp)
			return
		}
		for i := range nums {
			var newNums []int
			newNums = append(newNums, nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)

			var newTmp []int
			newTmp = append(newTmp, tmp...)
			newTmp = append(newTmp, nums[i])

			backtrack(newNums, newTmp)
		}
	}

	backtrack(nums, nil)
	return res
}
