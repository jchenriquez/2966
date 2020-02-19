package threesum

func threeSum(sums []int) (result [][]int) {
	negatedMp := make(map[int]map[int]bool)

	for index, sum := range sums {
		st, _ := negatedMp[sum*-1]
		st[index] = true
	}

	for i := 0; i < len(sums); i++ {

		valAtIndex := sums[i]
		st, _ := negatedMp[valAtIndex*-1]

		delete(st, i)

		for j := 0; j < len(sums); j++ {
			if j == i {
				continue
			}

			sm := sums[i] + sums[j]

			stTmp, _ := negatedMp[sm]

			if len(stTmp) != 0 {
				for ind := range stTmp {
					result = append(result, []int{i, j, ind})
				}
			}

		}

		st[i] = true
	}

	return
}
