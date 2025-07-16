package twosum

func twoSum(nums []int, target int) []int {
	hmap := make(map[int]int)

	for i, v := range nums {
		rem := target - v
		j, k := hmap[rem]
		if k {
			return []int{j, i}
		}

		hmap[v] = i
	}

	return []int{0, 0}
}
