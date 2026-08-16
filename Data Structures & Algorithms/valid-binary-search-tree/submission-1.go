/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func valid(root *TreeNode,l int, r int) bool{
	if root == nil{
		return true
	}
	if root.Val <= l || root.Val >= r{
		return false
	}
	return valid(root.Left, l, root.Val) && valid(root.Right, root.Val, r)
	
}
func isValidBST(root *TreeNode) bool {
    return valid(root,math.MinInt, math.MaxInt)
}
