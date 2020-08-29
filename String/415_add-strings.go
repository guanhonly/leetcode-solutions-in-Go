package String

import "strconv"

func addStrings(num1 string, num2 string) string {
	carry := 0
	var res string
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		currSum := x + y + carry
		if currSum >= 10 {
			res = strconv.Itoa(currSum-10) + res
			carry = 1
		} else {
			res = strconv.Itoa(currSum) + res
			carry = 0
		}
	}
	return res
}
