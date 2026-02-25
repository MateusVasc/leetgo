package findallnumbersdisappearedinanarray

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if v < 0 {
			v = -v
		}

		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}

	sl := []int{}

	for i, v := range nums {
		if v > 0 {
			sl = append(sl, i+1)
		}
	}

	return sl
}
