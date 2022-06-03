package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodes := []*ListNode{}
	node := head

	for node != nil {
		nodes = append(nodes, node)
		node = node.Next
	}
	return nodes[(len(nodes))/2]
}
