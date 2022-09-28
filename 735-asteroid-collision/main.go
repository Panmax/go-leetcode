package main

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		alive := true
		for alive && asteroid < 0 && len(stack) > 0 && stack[len(stack)-1] > 0 {
			alive = stack[len(stack)-1] < -asteroid
			if stack[len(stack)-1] <= -asteroid {
				stack = stack[:len(stack)-1]
			}
		}
		if alive {
			stack = append(stack, asteroid)
		}
	}

	return stack
}
