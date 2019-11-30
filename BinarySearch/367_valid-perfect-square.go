/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/valid-perfect-square/
 */

package BinarySearch

func isPerfectSquare(num int) bool {
    if num == 0 || num == 1 {
        return true
    }
    left, right := 0, num
    for left <= right {
        mid := (left+right)/2
        power := mid * mid
        if power < num {
            left = mid + 1
        } else if power > num {
            right = mid-1
        } else {
            return true
        }
    }
    return false
}

func isPerfectSquare2(num int) bool {
    odd, power := 1, 0
    for {
        power += odd
        odd += 2
        if power == num {
            return true
        }
        if power > num {
            return false
        }
    }
    return false
}