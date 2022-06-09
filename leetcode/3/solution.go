package main

func lengthOfLongestSubstring(s string) int {
	chars := map[rune]bool{}
	maxLen, left := 0, 0

	for i, char := range s {
		for chars[char] == true {
			chars[rune(s[left])] = false
			left++
		}
		chars[char] = true

		if maxLen < i-left+1 {
			maxLen = i - left + 1
		}
	}
	return maxLen
}
