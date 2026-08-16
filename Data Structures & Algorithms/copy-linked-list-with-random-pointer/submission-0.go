/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil{
		return nil
	}
    m := make(map[*Node]*Node) 
	cur := head
	for cur !=nil{
		c := &Node{ Val: cur.Val}
		m[cur]=c
		cur = cur.Next
	}
	cur = head
	for cur != nil{
		copy := m[cur]
		copy.Next = m[cur.Next]
		copy.Random = m[cur.Random]
		cur = cur.Next
	}
	return m[head]
}
