/**
 * Difficulty: Medium
 * Question link: https://leetcode.com/problems/counting-bits/
 */
package BitwiseOperations

func countBits(num int) []int {
    bits := make([]int, num+1)
    for i:=1; i<=num; i++ {
        bits[i] = bits[i&(i-1)] + 1
    }
    return bits
}