package main

func majorityElement_20221002(nums []int) []int {
	num1, num2 := nums[0], nums[0]
	var count1, count2 int
	for _, num := range nums {
		if num1 == num {
			count1++
			continue
		}
		if num2 == num {
			count2++
			continue
		}
		if count1 == 0 {
			num1 = num
			count1++
			continue
		}
		if count2 == 0 {
			num2 = num
			count2++
			continue
		}
		count1--
		count2--
	}
	count1, count2 = 0, 0
	for _, num := range nums {
		if num1 == num {
			count1++
		} else if num2 == num {
			count2++
		}
	}

	res := make([]int, 0)
	if count1 > len(nums)/3 {
		res = append(res, num1)
	}
	if count2 > len(nums)/3 {
		res = append(res, num2)
	}
	return res
}
