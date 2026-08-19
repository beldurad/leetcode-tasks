type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type travElem struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)
	queue := make([]travElem, 0)
	if root != nil {
		queue = append(queue, travElem{node: root, level: 0})
	}
	for len(queue) != 0 {
		cur := queue[0]
		if len(result) <= cur.level {
			result = append(result, make([]int, 0))
		}
		result[cur.level] = append(result[cur.level], cur.node.Val)
		if cur.node.Left != nil {
			queue = append(queue, travElem{node: cur.node.Left, level: cur.level + 1})
		}
		if cur.node.Right != nil {
			queue = append(queue, travElem{node: cur.node.Right, level: cur.level + 1})
		}

		queue = queue[1:]
	}
	return result
}
