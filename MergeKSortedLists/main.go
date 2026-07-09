import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type minHeap struct {
	internal []int
	size     int
}

func (h *minHeap) push(val int) {
	slice := append(h.internal, val)
	h.size += 1
	valIndex := h.size - 1
	parentIndex := (valIndex - 1) / 2
	for parentIndex >= 0 {
		if slice[valIndex] >= slice[parentIndex] {
			break
		}
		slice[valIndex], slice[parentIndex] = slice[parentIndex], slice[valIndex]
		valIndex = parentIndex
		parentIndex = (valIndex - 1) / 2
	}
	h.internal = slice
}

func (h *minHeap) pop() (int, error) {
	slice := h.internal
	size := h.size
	if size == 0 {
		return 0, fmt.Errorf("heap is empty")
	}
	res := slice[0]
	slice[0], slice[size-1] = slice[size-1], 0
	size -= 1
	valIndex := 0
	nextIndex := valIndex*2 + 1
	for nextIndex < size {
		minIndex := nextIndex
		if nextIndex+1 < size && slice[nextIndex] > slice[nextIndex+1] {
			minIndex += 1
		}
		if slice[valIndex] <= slice[minIndex] {
			break
		}
		slice[valIndex], slice[minIndex] = slice[minIndex], slice[valIndex]
		valIndex = minIndex
		nextIndex = valIndex*2 + 1
	}
	h.internal = slice
	h.size = size
	return res, nil
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	minHeap := &minHeap{
		internal: make([]int, 0),
	}

	for _, list := range lists {
		cur := list
		for cur != nil {
			minHeap.push(cur.Val)
			cur = cur.Next
		}
	}
	fmt.Println(minHeap.internal)
	var root *ListNode
	var cur *ListNode
	elem, err := minHeap.pop()
	if err != nil {
		return nil
	}
	root = &ListNode{
		Val: elem,
	}
	cur = root
	for elem, err = minHeap.pop(); err == nil; elem, err = minHeap.pop() {
		cur.Next = &ListNode{
			Val: elem,
		}
		cur = cur.Next
	}
	return root
}
