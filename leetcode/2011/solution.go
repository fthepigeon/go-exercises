package main

func finalValueAfterOperations(operations []string) int {
	val := 0

	for _, op := range operations {
		switch op[1] {
		case '-':
			val--
		case '+':
			val++
		}
	}
	return val
}
