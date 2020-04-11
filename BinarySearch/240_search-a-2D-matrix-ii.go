package BinarySearch

func searchMatrixII(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    if n == 0 {
        return false
    }
    row, col := 0, n-1
    for row < m && col >=0 {
        if matrix[row][col] == target {
            return true
        } else if matrix[row][col] < target {
            row++
        } else {
            col--
        }
    }
    return false
}