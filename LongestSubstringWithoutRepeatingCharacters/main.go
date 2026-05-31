func lengthOfLongestSubstring(s string) int {
	//Map, which sets each character it's index in string
	curCharsMap := make(map[rune]int)
	maxLen := 0
	curLen := 0
	curStartIndex := 0
	for i, c := range s {
		if prevIndex, ok := curCharsMap[c]; !ok || prevIndex < curStartIndex {
			curLen += 1
		} else {
			if curLen > maxLen {
				maxLen = curLen
			}
			curLen -= (prevIndex - curStartIndex)
			curStartIndex = prevIndex + 1
		}
		curCharsMap[c] = i
	}
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}