package main

func intToRoman_20230517(num int) string {
	var res string
	for _, symbol := range valueSymbols {
		for num >= symbol.value {
			num -= symbol.value
			res += symbol.symbol
		}
		if num == 0 {
			break
		}
	}
	return res
}
