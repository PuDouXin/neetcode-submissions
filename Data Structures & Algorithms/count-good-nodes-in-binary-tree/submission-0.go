/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 func dfs(root *TreeNode, maxVal int)int{
	if root == nil{
		return 0
	}
	res :=0
	if root.Val >=maxVal{
		maxVal = root.Val
		res =1
	}
	res += dfs(root.Left, maxVal)
	res += dfs(root.Right, maxVal)
	return res
 }

func goodNodes(root *TreeNode) int {
    if root == nil{
		return 0
	}
	return dfs(root, root.Val)

}
