package main

func intToRoman_20230331(num int) string {
	var res string
	for _, value := range valueSymbols {
		for num >= value.value {
			num -= value.value
			res += value.symbol
		}
		if num == 0 {
			break
		}
	}
	return res
}
