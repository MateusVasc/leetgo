package palindromenumber

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	n := x
	out := 0

	for n > 9 {
		out = out*10 + n%10
		n /= 10
	}

	out = out*10 + n
	return out == x
}
