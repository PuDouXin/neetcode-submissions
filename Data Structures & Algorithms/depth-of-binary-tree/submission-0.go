/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func depth(root *TreeNode, l int) int{
	for root.Left ==nil && root.Right ==nil{
		return l
	}
	var leftL, rightL int
	if root.Left !=nil{
		leftL = depth(root.Left,l+1)
	}
	if root.Right != nil{
		rightL = depth(root.Right,l+1)
	}
	if leftL>=rightL{
		return leftL
	}
	return rightL
}
func maxDepth(root *TreeNode) int {
    if root == nil{
		return 0
	}
	return depth(root,1)
}
