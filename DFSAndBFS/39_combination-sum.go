package DFSAndBFS

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	if len(candidates) == 0 {
		return res
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	var combination []int
	var helper func(startIndex, target int)
	helper = func(startIndex, target int) {
		if target == 0 {
			res = append(res, append([]int{}, combination...))
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			if target < candidates[i] {
				continue
			}
			if i > 0 && candidates[i] == candidates[i-1] {
				continue
			}
			combination = append(combination, candidates[i])
			helper(i, target-candidates[i])
			combination = combination[:len(combination)-1]
		}
	}
	helper(0, target)
	return res
}
