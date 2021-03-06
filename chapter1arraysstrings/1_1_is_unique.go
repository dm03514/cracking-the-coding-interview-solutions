package chapter1arraysstrings

// IsUnique checks to see if a string contains all unique characters.
//
// Assumes: ASCII input and dictionary size of 128.
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func IsUnique(s string) bool {
	if len(s) > 128 {
		return false
	}

	seen := make(map[uint8]struct{})
	for i := 0; i < len(s); i++ {
		char := s[i]
		if _, ok := seen[char]; ok {
			return false
		}

		seen[char] = struct{}{}
	}

	return true
}
