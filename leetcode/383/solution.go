package main

func canConstruct(ransomNote string, magazine string) bool {
	counter := map[rune]int{}

	for _, c := range magazine {
		counter[c]++
	}

	for _, c := range ransomNote {
		if counter[c] == 0 {
			return false
		}
		counter[c]--
	}
	return true
}
