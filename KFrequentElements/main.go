type elem struct {
	number    int
	frequency int
}

func topKFrequent(nums []int, k int) []int {
	elemSlice := make([]*elem, 0)
	freqMap := make(map[int]*elem)
	for _, n := range nums {
		e, ok := freqMap[n]
		if !ok {
			newElem := &elem{
				number:    n,
				frequency: 1,
			}
			freqMap[n] = newElem
			elemSlice = append(elemSlice, newElem)
			continue
		}
		e.frequency += 1
	}
	sort.Slice(elemSlice, func(i, j int) bool {
		return elemSlice[i].frequency >= elemSlice[j].frequency
	})
	res := make([]int, k)
	for i := 0; i < k; i += 1 {
		res[i] = elemSlice[i].number
	}
	return res
}
