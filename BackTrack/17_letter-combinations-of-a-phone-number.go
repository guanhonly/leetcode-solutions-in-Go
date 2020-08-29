package BackTrack

var mapping = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	combinations = []string{}
	backtrack(0, digits, "")
	return combinations
}

func backtrack(index int, digits, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		digit := digits[index]
		letters := mapping[digit]
		for i := 0; i < len(letters); i++ {
			backtrack(index+1, digits, combination+string(letters[i]))
		}
	}
}
