package threesum

import (
	"fmt"
	"sort"
)

// ThreeSum will return the triplets that sum to zero
func ThreeSum(sums []int) (result [][]int) {
	negatedMp := make(map[int]map[int]bool)
	ansSet := make(map[string]bool)
	for index, sum := range sums {
		st, in := negatedMp[sum*-1]

		if !in {
			st = make(map[int]bool)
		}

		st[index] = true

		negatedMp[sum*-1] = st
	}

	if len(negatedMp) == 1 {
		indeces, in := negatedMp[0]
		if in && len(indeces) >= 3 {
			return [][]int{{0, 0, 0}}
		}
		return [][]int{}
	}

	for i := 0; i < len(sums); i++ {

		valAtIndex := sums[i]
		st, _ := negatedMp[valAtIndex*-1]

		delete(st, i)

		for j := i + 1; j < len(sums); j++ {

			sm := sums[i] + sums[j]
			jSt, _ := negatedMp[sums[j]*-1]
			delete(jSt, j)

			stTmp, in := negatedMp[sm]

			if in && len(stTmp) > 0 {
				for ind := range stTmp {
					ans := []int{sums[i], sums[j], sums[ind]}
					sort.Ints(ans)
					inStr := fmt.Sprintf("%d%d%d", ans[0], ans[1], ans[2])

					if _, in := ansSet[inStr]; !in {
						result = append(result, ans)
						ansSet[inStr] = true
					}
					break
				}
			}

			jSt[j] = true

		}

		st[i] = true
	}

	return
}
