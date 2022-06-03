package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	values := []int{}
	node := head

	for node != nil {
		values = append(values, node.Val)
		node = node.Next
	}

	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-1-i] {
			return false
		}
	}
	return true
}
