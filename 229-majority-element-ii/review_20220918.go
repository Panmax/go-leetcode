package main

func majorityElement_20220918(nums []int) []int {
	res := make([]int, 0)
	if len(nums) == 0 {
		return res
	}
	nums1, nums2 := nums[0], nums[0]
	counter1, counter2 := 0, 0
	for _, num := range nums {
		if nums1 == num {
			counter1++
			continue
		}
		if nums2 == num {
			counter2++
			continue
		}
		if counter1 == 0 {
			nums1 = num
			counter1++
			continue
		}
		if counter2 == 0 {
			nums2 = num
			counter2++
			continue
		}
		counter1--
		counter2--
	}
	counter1, counter2 = 0, 0
	for _, num := range nums {
		if nums1 == num {
			counter1++
		} else if nums2 == num {
			counter2++
		}
	}
	if counter1 > len(nums)/3 {
		res = append(res, nums1)
	}
	if counter2 > len(nums)/3 {
		res = append(res, nums2)
	}
	return res
}
