package main

import "sort"

func minimumSum(num int) int {
	var digits []int

	for i := 0; i < 4; i++ {
		digits = append(digits, num%10)
		num /= 10
	}
	sort.Ints(digits)
	new1 := digits[0]*10 + digits[2]
	new2 := digits[1]*10 + digits[3]

	return new1 + new2
}
