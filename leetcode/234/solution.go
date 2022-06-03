package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	values := map[int]int{}
	node := head
	index := 0

	for {
		values[index] = node.Val

		if node.Next == nil {
			break
		}
		node = node.Next
		index++
	}

	for i := 0; i < len(values)/2; i++ {
		if values[i] != values[len(values)-1-i] {
			return false
		}
	}
	return true
}
