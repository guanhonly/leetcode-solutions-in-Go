package DivideAndConquer

import "math"

/**
 * Time complexity: O(n)
 * Space complexity: O(1)
 */
func maxSubArray(nums []int) int {
    maxSum := 0
    finalMaxSum := math.MinInt32
    for _, num := range nums {
        maxSum += num
        if maxSum < num {
            maxSum = num
        }
        if maxSum > finalMaxSum {
            finalMaxSum = maxSum
        }
    }
    return finalMaxSum
}

/**
 * Divide and conquer.
 * Divide the question to three sub-question: left, middle and right.
 * First, calculate the left and right, then for the array in middle
 * that intersect with left and right, calculate the middle sum, finally
 * pick the maximum of them.
 * Time complexity: O(nlogn)
 * Space complexity: O(1)
 */
func maxSubArray2(nums []int) int {
    n := len(nums)
    if n==1 {
        return nums[0]
    }
    maxLeft := maxSubArray2(nums[:n/2])
    maxRight := maxSubArray2(nums[n/2:])
    maxL := nums[n/2 - 1]
    tmp := 0
    for i:= n/2-1; i >=0; i-- {
        tmp += nums[i]
        maxL = max(maxL, tmp)
    }
    maxR := nums[n/2]
    tmp = 0
    for i:=n/2; i<n; i++ {
        tmp += nums[i]
        maxR = max(maxR, tmp)
    }
    return max(maxLeft, maxRight, maxL+maxR)
}

func max(nums... int) int {
    res := nums[0]
    for _, num := range nums {
        if num > res {
            res = num
        }
    }
    return res
}