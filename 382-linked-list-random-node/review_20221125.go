package main

import "math/rand"

func (s *Solution) GetRandom_20221125() int {
	var ret int

	for node, i := s.head, 1; node != nil; node = node.Next {
		if rand.Intn(i) == 0 {
			ret = node.Val
		}
		i++
	}
	return ret
}
