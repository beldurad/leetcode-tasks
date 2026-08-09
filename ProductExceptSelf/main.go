func productExceptSelf(nums []int) []int {
	answer := make([]int, len(nums))
	answer[0] = 1
	for i := 1; i < len(answer); i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		prev := 1
		if i+2 < len(nums) {
			prev = nums[i+2]
		}
		nums[i+1] *= prev
		answer[i] *= nums[i+1]
	}
	return answer
}
