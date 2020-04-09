package BinarySearch

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[right] {
			if target >= nums[left] && target <= nums[mid] {
				right = mid
			} else {
				left = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
