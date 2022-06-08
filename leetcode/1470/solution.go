package main

func shuffle(nums []int, n int) []int {
	ret := make([]int, 2*n)

	for i := 0; i < n; i++ {
		ret[i*2] = nums[i]
		ret[i*2+1] = nums[i+n]
	}
	return ret
}
