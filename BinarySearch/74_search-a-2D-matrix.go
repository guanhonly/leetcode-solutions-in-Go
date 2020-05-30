package BinarySearch

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}
	for i := 0; i < m; i++ {
		if target > matrix[i][n-1] {
			continue
		}
		left, right := 0, n-1
		for left < right {
			mid := left + (right-left)/2
			if matrix[i][mid] == target {
				return true
			}
			if matrix[i][mid] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		return matrix[i][left] == target
	}
	return false
}
