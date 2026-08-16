/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func merge2List(a *ListNode, b *ListNode) *ListNode{
	dummy := &ListNode{}
	res := dummy
	for a!=nil && b != nil{
		if a.Val <= b.Val{
			res.Next = a
			a = a.Next
		}else{
			res.Next = b
			b = b.Next
		}
		res =res.Next
	}
	for a !=nil{
		res.Next = a
		a = a.Next
		res =res.Next
	}
	for b != nil{
		res.Next = b
		b = b.Next
		res =res.Next
	}
	return dummy.Next
}
func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists)==0{
		return nil
	}
	for i:=1;i<len(lists);i++{
		lists[i] = merge2List(lists[i-1],lists[i])
	}
	return lists[len(lists)-1]
}
