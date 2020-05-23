package TwoPointers

func moveZeroes(nums []int) {
	for left, right := 0, 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[right], nums[left] = nums[left], nums[right]
			left++
		}
	}
}
