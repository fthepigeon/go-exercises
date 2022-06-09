package main

func lengthOfLongestSubstring(s string) int {
	if sLen := len(s); sLen == 0 || sLen == 1 {
		return sLen
	}

	chars := map[rune]bool{}
	maxLen, left := 0, 0

	for i, char := range s {
		for chars[char] == true {
			chars[rune(s[left])] = false
			left++
		}
		chars[char] = true

		if maxLen < i-left {
			maxLen = i - left
		}
	}
	return maxLen + 1
}
