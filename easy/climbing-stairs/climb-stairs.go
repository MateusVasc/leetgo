package climbingstairs

func climbStairs(n int) int {
	if n < 3 {
		return n
	}

	lenStairs := 2
	step1 := 1
	step2 := 2

	for lenStairs < n {
		next := step2 + step1
		step1 = step2
		step2 = next
		lenStairs++
	}

	return step2
}

// leetcode.com/problems/climbing-stairs/
