package main

func countBits(n int) []int {
	out := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		out[i] = hammingWeight(i)
	}
	return out
}

func hammingWeight(num int) int {
	ones := 0

	for num > 0 {
		ones += num & 1
		num >>= 1
	}
	return ones
}
