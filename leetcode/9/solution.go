package main

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	if x <= 10 {
		return true
	}

	rev := 0
	n := x

	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}

	return rev == x
}
