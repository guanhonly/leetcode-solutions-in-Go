package DFSAndBFS

import "bytes"

func solveNQueens(n int) [][]string {
	var res [][]string
	if n <= 0 {
		return res
	}
	search(&[]int{}, &res, n)
	return res
}

// In rows, every element refers to the index of the Queen in a row.
func search(rows *[]int, res *[][]string, n int) {
	if len(*rows) == n {
		*res = append(*res, drawBoard(*rows))
		return
	}
	for i := 0; i < n; i++ {
		if !isValid(*rows, i) {
			continue
		}
		*rows = append(*rows, i)
		search(rows, res, n)
		*rows = (*rows)[:len(*rows)-1]
	}
}

func drawBoard(rows []int) []string {
	n := len(rows)
	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		var row bytes.Buffer
		for j := 0; j < n; j++ {
			if rows[i] == j {
				row.WriteString("Q")
			} else {
				row.WriteString(".")
			}
		}
		res = append(res, row.String())
	}
	return res
}

func isValid(rows []int, newColIndex int) bool {
	rowIndex := len(rows)
	for i := 0; i < len(rows); i++ {
		if rows[i] == newColIndex {
			return false
		}
		if i+rows[i] == rowIndex+newColIndex {
			return false
		}
		if i-rows[i] == rowIndex-newColIndex {
			return false
		}
	}
	return true
}
