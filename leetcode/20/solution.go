package main

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	queue := []rune{}
	brackets := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, br := range s {
		// opening bracket
		if cbr, ok := brackets[br]; ok {
			queue = append(queue, cbr)
			continue
		}

		// closing bracket
		if len(queue) == 0 || br != queue[len(queue)-1] {
			return false // current bracket doesn't close the last one
		}
		queue = queue[:len(queue)-1]
	}
	return len(queue) == 0
}
