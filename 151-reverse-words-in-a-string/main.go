package main

func reverseWords(s string) string {
	// 去除左右空格
	left, right := 0, len(s)-1
	for left <= right && s[left] == ' ' {
		left++
	}
	for left <= right && s[right] == ' ' {
		right--
	}

	// 放入 bytes 数组中，并去除中间多余空格
	bytes := make([]byte, 0)
	s = s[left : right+1]
	for i := range s {
		if s[i] != ' ' {
			bytes = append(bytes, s[i])
		} else if len(bytes) != 0 && bytes[len(bytes)-1] != ' ' {
			bytes = append(bytes, s[i])
		}
	}
	// 所有字符翻转
	reverse(bytes, 0, len(bytes))

	// 按单词翻转
	start := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			reverse(bytes, start, i)
			start = i + 1
		}
	}
	// 翻转最后一个单词
	reverse(bytes, start, len(bytes))
	return string(bytes)
}

func reverse(bytes []byte, start, end int) {
	for i := start; i < (start+end)/2; i++ {
		bytes[i], bytes[(start+end)-i-1] = bytes[(start+end)-i-1], bytes[i]
	}
}

func main() {
	print(reverseWords(" world hello "))
}
