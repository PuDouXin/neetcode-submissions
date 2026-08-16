/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    nodeMap := make(map[*Node]*Node)
	var dfs func(*Node)*Node
	dfs = func(node *Node)*Node{
		if node == nil{
			return nil
		}
		if node, exist:=nodeMap[node]; exist{
			return node
		}
		newNode := &Node{Val:node.Val}
		nodeMap[node] = newNode
		for _,neighbor := range node.Neighbors{
			newNode.Neighbors =append(newNode.Neighbors, dfs(neighbor))
		}
		return newNode

	}
	return dfs(node)
	
}
