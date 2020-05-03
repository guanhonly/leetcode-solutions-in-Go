package DFSAndBFS

func permute(nums []int) [][]int {
	n := len(nums)
	var res [][]int
	if n == 0 {
		return res
	}
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
			if !visited[i] {
				visited[i] = true
				permutation = append(permutation, nums[i])
				helper()
				permutation = permutation[:len(permutation)-1]
				visited[i] = false
			}
		}
	}
	helper()
	return res
}
