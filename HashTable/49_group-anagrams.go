/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/group-anagrams/
 */
package HashTable

import "sort"

/**
 * The most intuitive way. The key of this question is to
 * find the hash function, which can generate unique hash key.
 * Time complexity: O(n*k*logk), n is the length of strs, k is
 * the length of the longest element in strs. Since sort is used,
 * must be multiplied by logk.
 * Space complexity: O(n*k), an extra hash table is used.
 */
func groupAnagrams(strs []string) [][]string {
    if len(strs) <=1 {
        return [][]string{strs}
    }
    hash := make(map[string][]string)
    for _, str := range strs {
        rStr := []rune(str)
        sort.Slice(rStr, func(i, j int) bool {
            return rStr[i] < rStr[j]
        })
        hash[string(rStr)] = append(hash[string(rStr)], str)
    }
    var res [][]string
    for key := range hash {
        res = append(res, hash[key])
    }
    return res
}

/**
 * This method uses a hash function that won't sort element, which
 * reduce the time complexity.
 * Time complexity: O(n*k)
 * Space complexity: O(n*k)
 */
func groupAnagrams2(strs []string) [][]string {
    hash := make(map[string]int)
    var result [][]string
    key := make([]rune, 26)
    for _, str := range strs {
        for _, r := range str {
            key[r-'a']++
        }
        strKey := string(key)
        if index, ok := hash[strKey]; ok {
            result[index] = append(result[index], str)
        } else {
            hash[strKey] = len(result)
            result = append(result, []string{str})
        }
    }
    return result
}

/**
 * A clever way, which use a hash function that generates numerical hash key.
 * Time complexity: O(n*k)
 * Space complexity: O(n*k)
 */
func groupAnagrams3(strs []string) [][]string {
    if len(strs) <=1 {
        return [][]string{strs}
    }
    hash := make(map[int32]int)
    var result [][]string
    for _, str := range strs {
        var mul int32 = 1
        var sum int32 = 0
        for _, r := range str {
            mul *= r
            sum += r
        }
        sum += mul
        if index, ok := hash[sum]; ok {
            result[index] = append(result[index], str)
        } else {
            hash[sum] = len(result)
            result = append(result, []string{str})
        }
    }
    return result
}

