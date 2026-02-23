package missingnumber

func missingNumber(nums []int) int {
	s := len(nums)
	fullSum := (s * (s + 1)) / 2
	missSum := 0

	for _, v := range nums {
		missSum += v
	}

	return fullSum - missSum
}
