func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	leftSideLen := (totalLen + 1) / 2
	// The last element of left side of nums1
	cur1 := len(nums1) / 2
	if len(nums1) == 0 {
		cur1 = -1
	}
	add := math.Max(float64(len(nums1)/4), 1)

	for {

		// Actual length of first array's left side should be less than expected total left side's length
		if cur1+1 > leftSideLen {
			cur1 -= int(add)
			add = math.Max(float64(add/2), 1)
			continue
		}

		// The last element of left side of nums2
		cur2 := leftSideLen - (cur1 + 1) - 1

		// Second array's left side must compensate for the remaining elements
		// This means that the length of the second array must be greater than or equal to the number of elements needed to build it on the left side.
		if cur2 >= len(nums2) {
			cur1 += int(add)
			add = math.Max(float64(add/2), 1)
			continue
		}

		if isInBounds(nums2, cur2) && isInBounds(nums1, cur1+1) && cur1 < len(nums1)-1 && nums2[cur2] > nums1[cur1+1] {
			cur1 += int(add)
			add = math.Max(float64(add/2), 1)
			continue
		}
		if isInBounds(nums1, cur1) && isInBounds(nums2, cur2+1) && cur2 < len(nums2)-1 && nums1[cur1] > nums2[cur2+1] {
			cur1 -= int(add)
			add = math.Max(float64(add/2), 1)
			continue
		}
		var leftMax = math.MinInt
		if isInBounds(nums1, cur1) {
			leftMax = nums1[cur1]
		}
		if isInBounds(nums2, cur2) && leftMax < nums2[cur2] {
			leftMax = nums2[cur2]
		}
		if totalLen%2 != 0 {
			return float64(leftMax)
		}
		rightMin := math.MaxInt
		if isInBounds(nums1, cur1+1) {
			rightMin = nums1[cur1+1]
		}
		if isInBounds(nums2, cur2+1) && rightMin > nums2[cur2+1] {
			rightMin = nums2[cur2+1]
		}
		return float64(leftMax+rightMin) / 2
	}
}

func isInBounds(slice []int, index int) bool {
	return index > -1 && index < len(slice)
}
