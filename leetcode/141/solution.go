package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fakeNode := &ListNode{}
	cNode := head

	for cNode != nil {
		if cNode.Next == fakeNode {
			return true
		}
		tmp := cNode.Next
		cNode.Next = fakeNode
		cNode = tmp

	}
	return false
}
