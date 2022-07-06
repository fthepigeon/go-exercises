package main

import "sort"

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	units := 0

	for i := 0; i < len(boxTypes); i++ {
		boxCount := boxTypes[i][0]

		if truckSize-boxCount >= 0 {
			units += boxCount * boxTypes[i][1]
			truckSize -= boxCount
		} else {
			units += truckSize * boxTypes[i][1]
			break
		}
	}
	return units
}
