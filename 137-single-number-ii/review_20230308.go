package main

func singleNumber_20230308(nums []int) int {
	var res int32
	for i := 0; i < 32; i++ {
		var b int
		for _, num := range nums {
			b += num >> i & 1
		}
		if b%3 == 1 {
			// |=
			res |= 1 << i
		}
	}
	return int(res)
}
