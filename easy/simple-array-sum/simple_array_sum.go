package simplearraysum

func SimpleArraySum(arr []int) int {
	out := 0

	for _, it := range arr {
		out += it
	}

	return out
}
