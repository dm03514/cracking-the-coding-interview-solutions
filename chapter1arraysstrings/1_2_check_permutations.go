package chapter1arraysstrings

// IsPermutation checks if either string is a permutation of the other.
//
// Assumes: ASCII input and dictionary size of 128.
//
// Time Complexity: O(n) - each string needs to be iterated over one time
// Space Complexity: O(s1) -
func IsPermutation(s1, s2 string) bool {
	// must be same length to be a permutation
	if len(s1) != len(s2) {
		return false
	}
	letterFrequency := make(map[uint8]int)
	for i := 0; i < len(s1); i++ {
		char := s1[i]
		letterFrequency[char]++
	}

	for i := 0; i < len(s2); i++ {
		char := s2[i]
		letterFrequency[char]--
		if letterFrequency[char] < 0 {
			return false
		}
	}

	return true
}
