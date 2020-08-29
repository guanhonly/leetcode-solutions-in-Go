package BinarySearch

func findMagicIndex(nums []int) int {
	length := len(nums)
	left, right := 0, length-1
	for left < right {
		mid := (left + right) / 2
		if (nums[mid] >= mid && nums[left] <= left) || (nums[mid] <= mid && nums[left] >= left) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if nums[left] == left {
		return left
	}
	return -1
}
