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
    queue := []*TreeNode{root}
	var res []string
	for len(queue)>0 {
		
				node := queue[0]
				queue = queue[1:]
				if node == nil{
					res = append(res,"N")
				}else{
				res = append(res, strconv.Itoa(node.Val))
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
				}
			
		}
	
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
     values := strings.Split(data,",")
	 if values[0] == "N"{
		return nil
	 }
	 rootVal, _ := strconv.Atoi(values[0])
	 root := &TreeNode{Val: rootVal}
	 queue := []*TreeNode{root}
	 index := 1
	 for len(queue) > 0 && index < len(values){
		node := queue[0]
		queue = queue[1:]
		if values[index] != "N"{
			left, _ := strconv.Atoi(values[index])
			node.Left = &TreeNode{Val: left}
			queue = append(queue, node.Left)
		}
		index++

		if index < len(values) && values[index] != "N" {
			right, _ := strconv.Atoi(values[index])
			node.Right = &TreeNode{Val: right}
			queue = append(queue, node.Right)
		}
		index++
	 }
	 return root
}
	 


