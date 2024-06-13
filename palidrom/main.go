package main

import "fmt"

func longestPalindromeSubstring(s string) string {
	if len(s) == 0 {
		return ""
	}

	n := len(s)
	dp := make([][]bool, n)
	start := 0
	maxLen := 1

	// Initialize the DP array
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	// Check for palindromes of length 2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start = i
			maxLen = 2
		}
	}

	// Check for palindromes of length greater than 2
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if dp[i+1][j-1] && s[i] == s[j] {
				dp[i][j] = true
				start = i
				maxLen = length
			}
		}
	}

	return s[start : start+maxLen]
}

func main() {
	input := "babad"
	result := longestPalindromeSubstring(input)
	fmt.Printf("Longest palindromic substring: %s\n", result)
}