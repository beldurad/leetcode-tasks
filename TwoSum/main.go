func twoSum(nums []int, target int) []int {
	intMap := make(map[int]int)
	for i := range nums {
		index, ok := intMap[nums[i]]
		if ok {
			res := make([]int, 2)
			res[0] = index
			res[1] = i
			return res
		}
		intMap[target-nums[i]] = i
	}
	return nil
}