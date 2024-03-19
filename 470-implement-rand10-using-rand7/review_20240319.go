package main

func rand10_20240319() int {
	for {
		r := (rand7()-1)*7 + (rand7() - 1)
		if r >= 1 && r <= 40 {
			return r%10 + 1
		}
	}
}
