package main

func containsDuplicate(nums []int) bool {
	counter := map[int]bool{}

	for _, num := range nums {
		if counter[num] {
			return true
		}
		counter[num] = true
	}
	return false
}
