package romantointeger

func RomanToInt(s string) int {
	sum := 0
	romanMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	i := 0
	for i < len(s)-1 {
		left := romanMap[string(s[i])]
		right := romanMap[string(s[i+1])]

		if left < right {
			sum += right - left
			i += 2
		} else {
			sum += left
			i++
		}
	}

	if i < len(s) {
		sum += romanMap[string(s[i])]
	}

	return sum
}
