package main

import "strings"

// first working solution, pretty inefficient
func convert(s string, numRows int) string {
	if numRows == 1 || numRows == len(s) || numRows > len(s) {
		return s
	}

	t := make([][]rune, numRows)

	for i := 0; i < numRows; i++ {
		t[i] = make([]rune, len(s)/2+1)
	}

	x, y, i := 0, 0, 0

outer:
	for {
		x = 0

		for j := 0; j < numRows; j++ {
			t[x][y] = rune(s[i])
			x++
			i++

			if i == len(s) {
				break outer
			}
		}

		x--

		for j := 0; j < numRows-2; j++ {
			y++
			x--
			t[x][y] = rune(s[i])
			i++

			if i == len(s) {
				break outer
			}
		}
		y++
	}

	var sb strings.Builder

	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s)/2+1; j++ {
			if t[i][j] != 0 {
				sb.WriteRune(t[i][j])
			}
		}
	}
	return sb.String()
}
