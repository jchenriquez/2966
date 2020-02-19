package threesum

func threeSum(sums []int) {
	negatedMp := make(map[int]map[int]bool)

	for index, sum := range sums {
		st, _ := negatedMp[sum*-1]
		st[index] = true
	}

	for i := 0; i < len(sums); i++ {

		for j := 0; j < len(sums); j++ {
			if j == i {
				continue
			}

		}
	}
}
