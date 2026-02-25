package findallnumbersdisappearedinanarray

func findDisappearedNumbers(nums []int) []int {
	s := len(nums)
	isIn := false
	sl := []int{}

	for i := range s {
		isIn = false
		i += 1
		for _, v := range nums {
			if i == v {
				isIn = true
			}
		}

		if !isIn {
			sl = append(sl, i)
		}
	}

	return sl
}
