package main

import "strings"

func mostWordsFound(sentences []string) int {
	max := 0

	for _, sntc := range sentences {
		if c := strings.Count(sntc, " "); c > max {
			max = c
		}
	}
	return max + 1
}
