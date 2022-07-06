package main

func fib(n int) int {
	if n == 0 {
		return 0
	}

	fib1, fib2 := 1, 1

	for i := 0; i < n-2; i++ {
		fib1, fib2 = fib2, fib1+fib2
	}

	return fib2
}
