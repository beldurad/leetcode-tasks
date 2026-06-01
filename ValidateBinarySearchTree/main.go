/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	return traverse(root, nil, nil)
}

func traverse(node *TreeNode, leftBorder, rightBorder *int) bool {
	if node == nil {
		return true
	}
	if leftBorder != nil && node.Val <= *leftBorder {
		return false
	}
	if rightBorder != nil && node.Val >= *rightBorder {
		return false
	}
	newBorder := &node.Val
	return traverse(node.Left, leftBorder, newBorder) && traverse(node.Right, newBorder, rightBorder)
}