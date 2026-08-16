/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {

	var res *ListNode
	for head != nil {
		n := head.Next
		head.Next = res
		res = head
		head = n
	}
	return res
}
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k <= 1 {
		return head
	}
	l := 1
	cur := head
	dummy := &ListNode{}
	res := dummy
	for cur != nil {
		if l == k {
			n := cur.Next
			cur.Next = nil
			tmp := reverse(head)

			for res.Next != nil {
				res = res.Next
			}
			res.Next = tmp

			head = n
			cur = n
			l = 1

		} else {
			cur = cur.Next
			l++
		}
	}
	if head != nil{
		for res.Next != nil {

			res = res.Next
		}
		res.Next = head
	}

	// if head != nil {
	// 	tmp := reverse(head)

	// 	for res.Next != nil {

	// 		res = res.Next
	// 	}
	// 	res.Next = tmp
	// }
	return dummy.Next
}