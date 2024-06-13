package main

import (
	"fmt"
)

func longestSubstringWithoutRepeating(s string) string {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0
	longestSubstring := ""

	for end := 0; end < len(s); end++ {
		if lastIndex, found := charIndex[s[end]]; found && lastIndex >= start {
			start = lastIndex + 1
		}

		charIndex[s[end]] = end
		currentSubstring := s[start : end+1]
		if len(currentSubstring) > maxLength {
			maxLength = len(currentSubstring)
			longestSubstring = currentSubstring
		}
	}

	return longestSubstring
}

func main() {
	input := "abcabcbb"
	result := longestSubstringWithoutRepeating(input)
	fmt.Printf("Longest substring without repeating characters: %s\n", result)
}
