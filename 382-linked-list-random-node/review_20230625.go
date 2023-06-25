package main

import "math/rand"

func (s *Solution) GetRandom_20230625() int {
	var ans int
	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
		i++
	}

	return ans
}
