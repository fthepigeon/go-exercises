package main

func sumSlice(s []int) int {
	sum := 0

	for _, v := range s {
		sum += v
	}
	return sum
}

func maximumWealth(accounts [][]int) int {
	max := 0

	for _, acc := range accounts {
		sum := sumSlice(acc)

		if sum > max {
			max = sum
		}
	}
	return max
}
