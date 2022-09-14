package main

// 超过 1/2 的最多有一个，超过 1/3 的最多有两个
func majorityElement(nums []int) []int {
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
			// 恢复到1
			counter1 = 1
			continue
		}
		if counter2 == 0 {
			num2 = num
			// 恢复到1
			counter2 = 1
			continue
		}
		counter1--
		counter2--
	}

	counter1, counter2 = 0, 0
	for _, num := range nums {
		if num1 == num {
			counter1++
		} else if num2 == num { // else if
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

// 放到字典里，把计数大于 t 的找出来
func majorityElementWithMap(nums []int) []int {
	res := make([]int, 0)
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for num, c := range m {
		if c > len(nums)/3 {
			res = append(res, num)
		}
	}
	return res
}
