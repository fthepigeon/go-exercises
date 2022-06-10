package main

import (
	"strconv"
	"strings"
)

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}

	var sb strings.Builder
	s := strconv.Itoa(x)

	for i := len(s); i > 0; i-- {
		sb.WriteByte(s[i-1])
	}
	return sb.String() == s
}
