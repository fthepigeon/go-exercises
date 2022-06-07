package main

func runningSum(nums []int) []int {
	runningSum := nums

	for i := 1; i < len(nums); i++ {
		runningSum[i] += runningSum[i-1]
	}
	return runningSum
}
