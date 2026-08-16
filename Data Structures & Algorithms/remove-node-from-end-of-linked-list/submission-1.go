/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var nodes []*ListNode
	cur := head
	for cur!= nil{
		nodes = append(nodes, cur)
		cur = cur.Next
	}
	if len(nodes)-n == 0{
		if len(nodes)==1{
			return nil
		}
		return head.Next
	}
	if len(nodes)-n+1 >=len(nodes){
		nodes[len(nodes)-n-1].Next = nil
	}else{
	nodes[len(nodes)-n-1].Next =nodes[len(nodes)-n+1]}
	return head
}
