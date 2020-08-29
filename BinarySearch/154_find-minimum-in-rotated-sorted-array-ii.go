package BinarySearch

func findMinII(nums []int) int {
	left, right := 0, len(nums)-1
	if nums[left] < nums[right] {
		return nums[left]
	}
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else if nums[mid] < nums[right] {
			right = mid
		} else { // duplicates exist
			right--
		}
	}
	return nums[left]
}
