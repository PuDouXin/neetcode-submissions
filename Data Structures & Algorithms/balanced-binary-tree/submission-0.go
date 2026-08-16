/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func isBalanced(root *TreeNode) bool {
    res:= true
	var dfs func(*TreeNode)int
	dfs = func (root *TreeNode) int{
	if root == nil{
		return 0
	}
	left := dfs(root.Left)
	right := dfs(root.Right)
	if left - right >1 || right -left >1{
		res = false
	}
	return 1+max(left, right)
}
	dfs(root)
return res
}
