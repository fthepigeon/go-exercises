package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}

	slow, fast := 0, 1

	for fast < len(nums) {
		if nums[fast] == nums[slow] {
			fast++
		} else {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
