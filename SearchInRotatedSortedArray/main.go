func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	first := nums[0]
	left, right := 0, len(nums)
	var bound int
	for left <= right {
		mid := left + (right-left)/2
		bound = rotateBound(nums, mid)
		if bound != -1 {
			break
		}
		if nums[mid] > first {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	left, right = 0, len(nums)
	for left <= right {
		mid := left + (right-left)/2
		rotateI := (mid + bound) % len(nums)
		if target == nums[rotateI] {
			return rotateI
		}
		if target > nums[rotateI] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func origValue(nums []int, k, i int) int {
	return nums[(i+k)%len(nums)]
}

func rotateBound(nums []int, i int) int {
	length := len(nums)
	left, right := i-1, i+1
	if right == length {
		right = 0
	}
	if left == -1 {
		left = length - 1
	}
	if nums[i] > nums[right] {
		return right
	}
	if nums[left] >= nums[i] {
		return i
	}
	return -1
}
