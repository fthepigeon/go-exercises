package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func calculateSum(v1, v2, rem int) (int, int) {
	sum := v1 + v2 + rem

	if sum/10 > 0 {
		return sum % 10, 1
	}
	return sum, 0
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n1, n2 := l1, l2
	v1, v2 := 0, 0
	rem := 0

	head := &ListNode{}
	cn := head

	for {
		if n1 != nil {
			v1 = n1.Val
			n1 = n1.Next
		} else {
			v1 = 0
		}

		if n2 != nil {
			v2 = n2.Val
			n2 = n2.Next
		} else {
			v2 = 0
		}

		cn.Val, rem = calculateSum(v1, v2, rem)

		if n1 == nil && n2 == nil {
			break
		}

		cn.Next = &ListNode{}
		cn = cn.Next
	}

	if rem > 0 {
		cn.Next = &ListNode{
			Val:  rem,
			Next: nil,
		}
	}
	return head
}
