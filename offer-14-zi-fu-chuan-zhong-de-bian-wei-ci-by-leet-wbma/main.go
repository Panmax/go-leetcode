package main

// https://leetcode.cn/problems/MPnaiL/submissions/
// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的某个变位词。
// 换句话说，第一个字符串的排列之一是第二个字符串的 子串 。
func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}
	s1Slot := [26]int{}
	s2Slot := [26]int{}
	for i := range s1 {
		s1Slot[s1[i]-'a']++
		s2Slot[s2[i]-'a']++
	}
	if s1Slot == s2Slot {
		return true
	}
	for i := s1Len; i < s2Len; i++ {
		s2Slot[s2[i]-'a']++
		s2Slot[s2[i-s1Len]-'a']--
		if s1Slot == s2Slot {
			return true
		}
	}
	return false
}
