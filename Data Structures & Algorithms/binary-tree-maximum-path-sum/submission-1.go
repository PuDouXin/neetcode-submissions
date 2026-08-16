/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func max(a int, b int)int{
	if a>=b{
		return a
	}
	return b
}
func maxPathSum(root *TreeNode) int {
    res := math.MinInt
	var dfs func(node *TreeNode)int
	dfs = func(node *TreeNode)int{
		if node == nil{
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		res = max(res, node.Val+max(left,0)+max(right,0))
		return node.Val+max(max(left,0), max(right,0))
	}
	t:=dfs(root)
	return max(res,t)
}
