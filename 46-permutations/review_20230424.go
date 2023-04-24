package main

func permute_20230424(nums []int) [][]int {
	var res [][]int
	var backtrace func([]int, []int)
	backtrace = func(nums []int, tmps []int) {
		if len(nums) == 0 {
			res = append(res, tmps)
			return
		}

		for i := 0; i < len(nums); i++ {
			var newNums []int
			newNums = append(newNums, nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)

			var newTmps []int
			newTmps = append(newTmps, tmps...)
			newTmps = append(newTmps, nums[i])
			backtrace(newNums, newTmps)
		}

	}

	backtrace(nums, nil)
	return res
}
