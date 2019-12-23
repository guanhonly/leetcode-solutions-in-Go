/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/permutation-in-string/
 */

package HashTable

func checkInclusion(s1 string, s2 string) bool {
    n := len(s1)
    if n > len(s2) {
        return false
    }
    charCountS1 := make([]int, 26)
    charCountS2 := make([]int, 26)
    for i:=0; i<n; i++ {
        charCountS1[s1[i]-'a']++
        charCountS2[s2[i]-'a']++
    }
    for i:=0; i<len(s2)-n; i++ {
        if matches(charCountS1, charCountS2) {
            return true
        }
        charCountS2[s2[i]-'a']--
        charCountS2[s2[i+n]-'a']++
    }
    return matches(charCountS1, charCountS2)
}

func matches(count1, count2 []int) bool {
    for i:=0; i<len(count1); i++ {
        if count1[i] != count2[i] {
            return false
        }
    }
    return true
}

func checkInclusion2(s1 string, s2 string) bool {
    n := len(s1)
    if n > len(s2) {
        return false
    }
    charCountS1 := make([]int, 26)
    charCountS2 := make([]int, 26)
    for i:=0; i<n; i++ {
        charCountS1[s1[i]-'a']++
        charCountS2[s2[i]-'a']++
    }
    count := 0
    for i:=0; i<26; i++ {
        if charCountS1[i] == charCountS2[i] {
            count++
        }
    }

    for i:=0; i<len(s2)-n; i++ {
        if count == 26 {
            return true
        }
        left, right := s2[i]-'a', s2[i+n]-'a'
        charCountS2[right]++
        if charCountS1[right] == charCountS2[right] {
            count++
        } else if charCountS1[right] == charCountS2[right]-1 {
            count--
        }
        charCountS2[left]--
        if charCountS1[left] == charCountS2[left] {
            count++
        } else if charCountS1[left] == charCountS2[left]+1 {
            count--
        }
    }
    return count == 26
}