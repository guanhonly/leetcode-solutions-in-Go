package TwoPointers

func sortColors(nums []int) {
	redIdx, blueIdx := 0, len(nums)-1
	for i := 0; i <= blueIdx; i++ {
		for nums[i] == 2 && i < blueIdx {
			nums[blueIdx], nums[i] = nums[i], nums[blueIdx]
			blueIdx--
		}
		for nums[i] == 0 && i > redIdx {
			nums[redIdx], nums[i] = nums[i], nums[redIdx]
			redIdx++
		}
	}
}
