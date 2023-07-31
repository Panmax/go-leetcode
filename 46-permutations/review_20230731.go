package main

func permute_20230731(nums []int) [][]int {
	var res [][]int
	var backtrace func([]int, []int)
	backtrace = func(nums []int, tmp []int) {
		if len(nums) == 0 {
			res = append(res, tmp)
		}
		for i := 0; i < len(nums); i++ {
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
