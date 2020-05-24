package HashTable

func subarraysDivByK(A []int, K int) int {
	hash := make([]int, K)
	hash[0] = 1
	sum, count := 0, 0
	for _, num := range A {
		sum = (sum + num) % K
		if sum < 0 {
			sum += K
		}
		count += hash[sum]
		hash[sum]++
	}
	return count
}
