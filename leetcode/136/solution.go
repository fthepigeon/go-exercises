package main

func singleNumber(nums []int) int {
	set := map[int]bool{}

	for _, n := range nums {
		if set[n] {
			delete(set, n)
		} else {
			set[n] = true
		}
	}
	for k := range set {
		return k
	}
	return -1
}
