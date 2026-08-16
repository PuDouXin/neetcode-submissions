/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func kthSmallest(root *TreeNode, k int) int {
    var dfs func(root *TreeNode)
	var nodes = []int{}
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		dfs(root.Left)
		nodes = append(nodes, root.Val)
		dfs(root.Right)

	}
	dfs(root)
	return nodes[k-1]
}
