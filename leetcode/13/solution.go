package main

func romanToInt(s string) int {
	var ret int
	roman := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i, c := range s {
		curr := roman[c]

		if i == len(s)-1 {
			return ret + curr
		}

		nxt := roman[rune(s[i+1])]

		if nxt > curr {
			ret -= curr
		} else {
			ret += curr
		}
	}
	return ret
}
