package mian

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	rows := make([]string, numRows)
	i := 0
	flag := -1
	for _, c := range s {
		rows[i] += string(c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	var res string
	for _, s := range rows {
		res += s
	}
	return res
}
