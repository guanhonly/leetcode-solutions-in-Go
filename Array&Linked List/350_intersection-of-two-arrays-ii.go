package LinkedList

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})
	index1 := 0
	index2 := 0
	index := 0
	m, n := len(nums1), len(nums2)
	intersectNums := make([]int, min(m, n))
	for index1 < m && index2 < n {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersectNums[index] = nums1[index1]
			index1++
			index2++
			index++
		}
	}
	return intersectNums[:index]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
