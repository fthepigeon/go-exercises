package main

func twoSum(nums []int, target int) []int {
	numbers := map[int]int{}

	for i, v := range nums {
		if idx, ok := numbers[target-v]; ok {
			return []int{i, idx}
		}
		numbers[v] = i
	}
	return []int{}
}
