package DivideAndConquer

func sortArray(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	tmp := make([]int, len(nums))
	mergeSort(0, len(nums)-1, nums, tmp)
	return nums
}

func mergeSort(start, end int, nums, tmp []int) {
	if start >= end {
		return
	}
	mid := start + (end-start)/2
	mergeSort(start, mid, nums, tmp)
	mergeSort(mid+1, end, nums, tmp)
	merge(start, end, nums, tmp)
}

func merge(start, end int, nums, tmp []int) {
	mid := start + (end-start)/2
	leftIndex := start
	rightIndex := mid + 1
	index := start

	for leftIndex <= mid && rightIndex <= end {
		if nums[leftIndex] < nums[rightIndex] {
			tmp[index] = nums[leftIndex]
			index++
			leftIndex++
		} else {
			tmp[index] = nums[rightIndex]
			index++
			rightIndex++
		}
	}

	for leftIndex <= mid {
		tmp[index] = nums[leftIndex]
		index++
		leftIndex++
	}

	for rightIndex <= end {
		tmp[index] = nums[rightIndex]
		index++
		rightIndex++
	}

	for i := start; i <= end; i++ {
		nums[i] = tmp[i]
	}
}

func sortArray2(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}
	quickSort(0, len(nums)-1, nums)
	return nums
}

func quickSort(start, end int, nums []int) {
	if start >= end {
		return
	}

	left := start
	right := end
	pivot := nums[(left+right)/2]
	for left <= right {
		for left <= right && nums[left] < pivot {
			left++
		}
		for left <= right && nums[right] > pivot {
			right--
		}
		if left <= right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	quickSort(start, right, nums)
	quickSort(left, end, nums)
}
