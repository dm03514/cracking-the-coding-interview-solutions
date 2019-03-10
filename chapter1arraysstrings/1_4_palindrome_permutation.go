package chapter1arraysstrings

import (
	"strings"
)

const space = 32

// IsPalindromePermutation checks to see if the string fulfills the rules for
// being a palindrome, and therefore supports a permutation which is a palindrome.
//
// Time Complexity: O(n)
// Space Complexity O(1) with respect to the alphabet size
func IsPalindromePermutation(s string) bool {
	lower := strings.ToLower(s)

	// get the letter frequencies
	letters := make(map[uint8]int)
	for i := 0; i < len(lower); i++ {
		// ignore spaces
		if lower[i] == space {
			continue
		}
		letters[lower[i]]++
	}

	// check that at most there is a single odd frequency
	foundOdd := false
	for _, v := range letters {
		if v%2 == 1 {
			if foundOdd {
				return false
			}
			foundOdd = true
		}
	}
	return true
}
