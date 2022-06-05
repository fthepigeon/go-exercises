package main

import "sort"

func kWeakestRows(mat [][]int, k int) []int {
	soldiers := [][]int{}
	out := make([]int, k)

	for i, row := range mat {
		sum := 0

		for _, x := range row {
			if x == 0 {
				break
			}
			sum += x
		}
		soldiers = append(soldiers, []int{sum, i})
	}
	sort.Slice(soldiers, func(i, j int) bool {
		if soldiers[i][0] != soldiers[j][0] {
			return soldiers[i][0] < soldiers[j][0]
		}
		return soldiers[i][1] < soldiers[j][1]
	})

	for i := 0; i < k; i++ {
		out[i] = soldiers[i][1]
	}
	return out
}
