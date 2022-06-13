package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var pNode *ListNode = nil
	cNode := head

	for cNode != nil {
		tmp := cNode.Next
		cNode.Next = pNode
		pNode = cNode
		cNode = tmp
	}
	return pNode
}
