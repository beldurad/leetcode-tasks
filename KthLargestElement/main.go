type slice []int

func (s slice) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s slice) Len() int {
	return len(s)
}

func (s slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func findKthLargest(nums []int, k int) int {
	sort.Sort(slice(nums))
	return nums[len(nums)-k]
}