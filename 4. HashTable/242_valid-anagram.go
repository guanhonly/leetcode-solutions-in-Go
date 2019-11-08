/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/valid-anagram/
 */
package HashTable

/**
 * Two maps are used, one is for s, another is for t.
 * Then compare if the two maps are equal.
 * Time complexity: O(n), n is the length of s
 * Space complexity: O(n)
 */
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hashS := make(map[int]int)
	hashT := make(map[int]int)
	for c := range s {
		hashS[c]++
		hashT[c]++
	}
	for key := range hashS {
		if hashS[key] != hashT[key] {
			return false
		}
	}
	return true
}

/**
 * Optimized way, only one array is used.
 * Time complexity: O(n)
 * Space complexity: O(1), since the length
 * of array is constant
 */
func isAnagram2(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    var hash [26]int
    for i:=0; i<len(s); i++ {
        hash[s[i] - 'a']++
        hash[t[i] - 'a']--
    }
    for _,i := range hash {
        if i != 0 {
            return false
        }
    }
    return true
}