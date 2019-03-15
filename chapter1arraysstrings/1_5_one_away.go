package chapter1arraysstrings

// IsOneAway checks two strings are a single edit from each other.
//
// Time Complexity O(n)
// Space Complexity O(1)
func IsOneAway(first, second string) bool {
	if len(first) == len(second) {
		return hasOneEditReplace(first, second)
	} else if len(first)+1 == len(second) {
		return hasOneEditInsert(first, second)
	} else if len(second)+1 == len(first) {
		return hasOneEditInsert(second, first)
	}
	return false
}

// hasOneEditInsert checks to see if each string
// is equal allowing for one missing character
func hasOneEditInsert(s1, s2 string) bool {
	i1 := 0
	i2 := 0
	for i1 < len(s1) && i2 < len(s2) {
		// check to see if the characters are equal
		if s1[i1] != s2[i2] {
			// if the indexes are no longer equal there has been more than
			// a single different character
			if i1 != i2 {
				return false
			}
			// allow for single different character by incrementing
			// the longer string (s2) skipping its character
			i2++
		} else {
			i1++
			i2++
		}
	}

	return true
}

// hasOneEditReplace checks each string
// allowing up to a single difference
func hasOneEditReplace(s1, s2 string) bool {
	hasDifference := false
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if hasDifference {
				return false
			}
			hasDifference = true
		}
	}
	return true
}
