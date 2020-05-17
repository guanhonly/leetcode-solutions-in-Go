package LinkedList

func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	intersect := make(map[int]bool)
	for _, num := range nums1 {
		set[num] = true
	}
	for _, num := range nums2 {
		if _, exists := set[num]; exists {
			intersect[num] = true
		}
	}
	var res []int
	for k := range intersect {
		res = append(res, k)
	}
	return res
}
