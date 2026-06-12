
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// bytes is a type representing an [TreeNode] Value converted to a slice of bytes
type bytes []byte

func zeroBytes() bytes {
	return bytes{1}
}
func (b bytes) isZero() bool {
	return len(b) >= 1 && b[0] == 1
}

func nilBytes() bytes {
	return bytes{0}
}
func (b bytes) isNil() bool {
	// a slice length less than one indicates that either the type is null or the type is invalid
	return len(b) <= 0 || b[0] == 0
}

func newBytes(n int) bytes {
	if n == 0 {
		return zeroBytes()
	}
	u := uint(n)
	b := make([]byte, 1)
	mask := uint(1<<8 - 1)

	// The least significant byte corresponds to the least significant index after 0
	for u > 0 {
		b = append(b, byte(mask&u))
		u >>= 8
	}

	//The first byte is informational, indicates the length of the integer in bytes
	b[0] = byte(len(b))
	return b
}

func (b bytes) toInt() int {
	res := uint(0)

	if b.isZero() {
		return 0
	}

	for i := len(b) - 1; i > 0; i -= 1 {
		res <<= 8
		mask := uint(b[i])
		res |= mask
	}
	return int(res)
}

func bytesFromString(s string) bytes {
	if s == "" {
		return zeroBytes()
	}
	bytesLen := s[0]
	if bytesLen == 0 {
		return nilBytes()
	}
	bound := int(math.Min(float64(bytesLen), float64(len(s))))
	return bytes(s[:bound])
}

type Codec struct {
}

func Constructor() Codec {

	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return string(nilBytes())
	}
	allBytes := make([]byte, 0)
	treeQueue := make([]*TreeNode, 1)
	treeQueue[0] = root

	// BFS
	for len(treeQueue) > 0 {
		if treeQueue[0] == nil {
			allBytes = append(allBytes, nilBytes()...)
			treeQueue = treeQueue[1:]
			continue
		}
		b := newBytes(treeQueue[0].Val)
		allBytes = append(allBytes, b...)
		treeQueue = append(treeQueue, treeQueue[0].Left, treeQueue[0].Right)
		treeQueue = treeQueue[1:]
	}
	return string(allBytes)
}

const (
	left  = 0
	right = 1
)

type parentQueueNode struct {
	node *TreeNode
	side int
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	parentQueue := make([]parentQueueNode, 0)
	rootBytes := bytesFromString(data)
	if rootBytes.isNil() {
		return nil
	}
	root := &TreeNode{
		Val: rootBytes.toInt(),
	}
	fmt.Printf("КОРЕНЬ: %d\n", root.Val)
	parentQueue = append(parentQueue, parentQueueNode{
		node: root,
		side: left,
	}, parentQueueNode{
		node: root,
		side: right,
	})
	bound := int(math.Min(float64(len(rootBytes)), float64(len(data))))
	data = data[bound:]
	for data != "" {
		curBytes := bytesFromString(data)
		if len(parentQueue) == 0 {
			panic("AAAAAAAAAAAAAAAAAAAAAAAA")
		}
		if curBytes.isNil() {
			fmt.Printf("Встретили нулевой элемент, он будет сыном %d\n", parentQueue[0].node.Val)
			data = data[1:]
			parentQueue = parentQueue[1:]
			continue
		}
		curNode := &TreeNode{
			Val: curBytes.toInt(),
		}
		if parentQueue[0].side == left {
			fmt.Printf("Встретили %d, присвоили как левого ребенка %d\n", curNode.Val, parentQueue[0].node.Val)
			parentQueue[0].node.Left = curNode
		} else {
			fmt.Printf("Встретили %d, присвоили как правого ребенка %d\n", curNode.Val, parentQueue[0].node.Val)
			parentQueue[0].node.Right = curNode
		}

		parentQueue = append(parentQueue, parentQueueNode{
			node: curNode,
			side: left,
		}, parentQueueNode{
			node: curNode,
			side: right,
		})
		parentQueue = parentQueue[1:]
		data = data[len(curBytes):]
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
