package main

import "math/rand"

func (s *Solution) GetRandom_20220812() int {
	res := 0
	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			res = node.Val
		}
		i++
	}
	return res
}
