package main

func majorityElement_20220916(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	num1, num2 := nums[0], nums[0]
	counter1, counter2 := 0, 0
	for _, num := range nums {
		if num1 == num {
			counter1++
			continue
		}
		if num2 == num {
			counter2++
			continue
		}
		if counter1 == 0 {
			num1 = num
			counter1++
			continue
		}
		if counter2 == 0 {
			num2 = num
			counter2++
			continue
		}
		counter1--
		counter2--
	}
	counter1, counter2 = 0, 0
	for _, num := range nums {
		if num == num1 {
			counter1++
		} else if num == num2 {
			counter2++
		}
	}
	if counter1 > len(nums)/3 {
		res = append(res, num1)
	}
	if counter2 > len(nums)/3 {
		res = append(res, num2)
	}
	return res
}
