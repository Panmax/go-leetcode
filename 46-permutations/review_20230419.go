package main

func permute_20230419(nums []int) [][]int {
	var res [][]int
	var backtrace func(nums, tmp []int)
	backtrace = func(nums, tmp []int) {
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

			backtrace(newNums, newTmp)
		}
	}

	backtrace(nums, nil)
	return res
}
