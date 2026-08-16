/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    var h *ListNode
	if list2 == nil && list1 == nil {
		return nil
	}
	if list1 == nil && list2 != nil{
		return list2
	}

	if list2 == nil && list1 != nil{
		return list1
	}
	if  list1.Val <= list2.Val{

		h = list1
		list1 = list1.Next
	} else {
		h = list2
		list2 = list2.Next
	}
	res := h

	for list1 != nil && list2 != nil {

		if list1.Val <= list2.Val {
			h.Next = list1
			list1 = list1.Next
		} else {
			h.Next = list2
			list2 = list2.Next
		}
		h = h.Next
	}

	for list1 != nil {
		h.Next = list1
		list1 = list1.Next
		h = h.Next
	}
	for list2 != nil {
		h.Next = list2
		list2 = list2.Next
		h = h.Next
	}

	return res
}
