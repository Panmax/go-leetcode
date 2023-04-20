package main

func singleNumber_20230420(nums []int) int {
	var ans int32
	for i := 0; i < 32; i++ {
		var total int
		for _, num := range nums {
			total += (num >> i) & 1
		}
		if total%3 == 1 {
			ans |= 1 << i
		}
	}

	return int(ans)
}
