/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func max(a, b int)int{
	if a>=b{
		return a
	}
	return b
}

func diameterOfBinaryTree(root *TreeNode) int {
	res:=0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int{
		if root == nil{
			return 0
		}

		left := dfs(root.Left)
		right := dfs(root.Right)

		res = max(res, left+right)
		return max(right, left)+1
	}
	dfs(root)
	return res
}
