package HashTable

func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, num := range nums {
		hash[num] = true
	}
	res := 0
	for k := range hash {
		if _, exists := hash[k-1]; !exists {
			currNum := k
			count := 1
			var exists bool
			_, exists = hash[currNum+1]
			for exists {
				currNum++
				count++
				_, exists = hash[currNum+1]
			}
			res = max(res, count)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
