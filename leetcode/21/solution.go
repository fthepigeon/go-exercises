package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	head := &ListNode{}
	cNode := head

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cNode.Next = list1
			list1 = list1.Next
		} else {
			cNode.Next = list2
			list2 = list2.Next
		}
		cNode = cNode.Next
	}

	if list2 == nil {
		cNode.Next = list1
	} else {
		cNode.Next = list2
	}
	return head.Next
}
