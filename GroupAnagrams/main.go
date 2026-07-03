func groupAnagrams(strs []string) [][]string {

	res := make([][]string, 0)

	groups := make(map[[26]int][]string)

	for _, s := range strs {
		chars := charCounts(s)
		if _, ok := groups[chars]; ok {
			groups[chars] = append(groups[chars], s)
			continue
		}
		newGroup := make([]string, 1)
		newGroup[0] = s
		groups[chars] = newGroup
	}
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}

func charCounts(s string) [26]int {
	var res [26]int
	for _, r := range s {
		res[int(r)-int('a')] += 1
	}
	return res
}
