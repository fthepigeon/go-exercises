package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	nodes := map[int]*ListNode{}
	node := head
	index := 0

	for {
		nodes[index] = node

		if node.Next == nil {
			break
		}

		node = node.Next
		index++
	}
	return nodes[(index+1)/2]
}
