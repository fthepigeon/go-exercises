package main

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	maxSeq := 0

	for _, n := range nums {
		if numSet[n] {
			continue
		}
		numSet[n] = true
	}

	for n := range numSet {
		if !numSet[n-1] {
			i := n

			for numSet[i+1] {
				i++
			}
			if i-n+1 > maxSeq {
				maxSeq = i - n + 1
			}
		}
	}
	return maxSeq
}
