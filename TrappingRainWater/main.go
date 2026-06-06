func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	left := 0
	right := 1

	result := 0

	secondMaxIndex := right

	for secondMaxIndex < len(height) {
		if right >= len(height) {
			for i := left + 1; i < secondMaxIndex; i += 1 {
				result += height[secondMaxIndex] - height[i]
			}
			left = secondMaxIndex
			right = secondMaxIndex + 1
			secondMaxIndex += 1
			continue
		}
		if height[right] >= height[left] {
			for i := left + 1; i < right; i += 1 {
				result += height[left] - height[i]
			}
			left = right
			right += 1
			secondMaxIndex = right
			continue
		}
		if height[right] > height[secondMaxIndex] {
			secondMaxIndex = right
		}
		right += 1
	}
	return result
}