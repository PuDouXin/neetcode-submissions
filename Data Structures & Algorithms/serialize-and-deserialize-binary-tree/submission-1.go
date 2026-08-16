/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil{
		return "N"
	}
    
	var res []string
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode){
		if root == nil{
			res = append(res, "N")
			return
		}
		res = append(res, strconv.Itoa(root.Val))
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
     values := strings.Split(data,",")
	 i := 0
	 var dfs func()*TreeNode
	dfs = func()*TreeNode{
		if values[i] == "N"{
			i++
			return nil
		}
		val, _ := strconv.Atoi(values[i])
		node := &TreeNode{Val:val}
		i++
		node.Left = dfs()
		node.Right = dfs()
		return node
	}

	 
	 return dfs()
}
	 


