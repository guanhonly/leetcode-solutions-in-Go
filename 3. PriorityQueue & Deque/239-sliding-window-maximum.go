/**
 * Difficulty: Hard
 * Question link: https://leetcode.com/problems/sliding-window-maximum/
 */

/**
 * Only slice is used indeed, but used like a deque.
 * Time complexity: O(n), n is the length of nums
 * Space complexity: O(n)
 */
func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    if n*k == 0 {
        return nil
    }
    if k == 1 {
        return nums
    }
    firstIndex := 0
    var indexWindow []int
    for i:=0; i<k; i++ {
        preClean(nums, &indexWindow, i, k)
        indexWindow = append(indexWindow, i)
        if nums[i] > nums[firstIndex] {
            firstIndex = i;
        }
    }
    res := make([]int, n-k+1)
    res[0] = nums[firstIndex]
    for i:=k; i<n; i++ {
        preClean(nums, &indexWindow, i, k)
        indexWindow = append(indexWindow, i)
        res[i-k+1] = nums[indexWindow[0]]
    }
    return res
}

func preClean(nums []int, indexWindow *[]int, i, k int) {
    if len(*indexWindow) > 0 && (*indexWindow)[0] == i-k {
        *indexWindow = (*indexWindow)[1:len(*indexWindow)]
    }
    for len(*indexWindow) > 0 && nums[i] > nums[(*indexWindow)[len(*indexWindow)-1]] {
        *indexWindow = (*indexWindow)[:len(*indexWindow)-1]
    }
}
