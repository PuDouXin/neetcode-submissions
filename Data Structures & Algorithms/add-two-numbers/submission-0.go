/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    n := 0
	res := &ListNode{Val: 0}
	result := res
	for l1 != nil && l2 != nil {
		r := l1.Val + l2.Val + n
		if r >= 10 {
			r = r - 10
			n = 1
		} else {
			n = 0
		}
		res.Next = &ListNode{Val: r}
		l1 = l1.Next
		l2 = l2.Next
		res = res.Next
	}
	for l1 != nil {
		r := l1.Val + n
		if r >= 10 {
			r = r - 10
			n = 1
		} else {
			n = 0
		}
		res.Next = &ListNode{Val: r}
		l1 = l1.Next
		res = res.Next
	}
	for l2 != nil {
		r := l2.Val + n
		if r >= 10 {
			r = r - 10
			n = 1
		} else {
			n = 0
		}
		res.Next = &ListNode{Val: r}
		l2 = l2.Next
		res = res.Next
	}
	if n == 1 {
		res.Next = &ListNode{Val: 1}
	}
	return result.Next


}
