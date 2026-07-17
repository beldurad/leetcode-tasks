func maxArea(height []int) int {
	res := 0
	left, right := 0, len(height)-1
	for left < right {
		x := right - left
		y := min(height[left], height[right])
		res = max(res, x*y)
		if height[left] > height[right] {
			right -= 1
		} else {
			left += 1
		}
	}
	return res

}
