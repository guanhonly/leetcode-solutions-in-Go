package DFSAndBFS

import "sort"

func permuteUnique(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	if n == 0 {
		return res
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	var permutation []int
	visited := make([]bool, n)
	var helper func()
	helper = func() {
		if len(permutation) == n {
			tmp := make([]int, n)
			copy(tmp, permutation)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if visited[i] || (i != 0 && nums[i-1] == nums[i] && !visited[i-1]) {
				continue
			}
			visited[i] = true
			permutation = append(permutation, nums[i])
			helper()
			permutation = permutation[:len(permutation)-1]
			visited[i] = false
		}
	}
	helper()
	return res
}
