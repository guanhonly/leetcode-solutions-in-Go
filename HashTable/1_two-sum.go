/**
 * Difficulty: Easy
 * Question link: https://leetcode.com/problems/two-sum/
 */
package HashTable
/**
 * It's the first question in leetcode, so I think it's not
 * necessary to explain more.
 * Time complexity: O(n), n is the length of nums.
 * Space complexity: O(n)
 */
func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		curNum := nums[i]
		if want, exist := hash[target-curNum]; exist {
			return []int{want, i}
		} else {
			hash[curNum] = i
		}
	}
	return nil
}
