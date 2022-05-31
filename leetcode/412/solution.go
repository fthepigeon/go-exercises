package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	ret := make([]string, n)

	for i := 1; i < n+1; i++ {
		if i%15 == 0 {
			ret[i-1] = "FizzBuzz"
			continue
		}
		if i%3 == 0 {
			ret[i-1] = "Fizz"
			continue
		}
		if i%5 == 0 {
			ret[i-1] = "Buzz"
			continue
		}
		ret[i-1] = strconv.Itoa(i)
	}
	return ret
}

// func main() {
// 	fmt.Println(fizzBuzz(20))
// }
