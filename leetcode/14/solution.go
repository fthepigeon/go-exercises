package main

func longestCommonPrefix(strs []string) string {
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return strs[0][:i]
			}
		}
	}
}
