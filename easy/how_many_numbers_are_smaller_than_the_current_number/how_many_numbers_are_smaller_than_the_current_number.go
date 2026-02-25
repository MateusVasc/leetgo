package howmanynumbersaresmallerthanthecurrentnumber

func smallerNumbersThanCurrent(nums []int) []int {
	sl := []int{}

	for _, v := range nums {
		c := 0

		for _, v2 := range nums {
			if v != v2 && v > v2 {
				c++
			}
		}

		sl = append(sl, c)
	}

	return sl
}
