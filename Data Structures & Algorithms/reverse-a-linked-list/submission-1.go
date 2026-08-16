/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var res *ListNode
	h := head
	for h != nil {
		temp := h.Next
		h.Next = res
		res = h
		h = temp
	}
	return res
}
