package BinarySearch

func twoSum(numbers []int, target int) []int {
	for i, num := range numbers {
		anotherIndex := binarySearch(numbers, target-num, i+1)
		if anotherIndex > 0 {
			return []int{i + 1, anotherIndex + 1}
		}
	}
	return []int{-1, -1}
}

func binarySearch(numbers []int, target, startIndex int) int {
	left, right := startIndex, len(numbers)-1
	for left <= right {
		mid := (left + right) / 2
		if numbers[mid] == target {
			return mid
		}
		if numbers[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left > right {
		return -1
	}
	return left
}
