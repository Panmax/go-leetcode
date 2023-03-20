package main

func intToRoman_20230320(num int) string {
	var res string
	for _, v := range valueSymbols {
		for num >= v.value {
			num -= v.value
			res += v.symbol
		}
		if num == 0 {
			break
		}
	}
	return res
}
