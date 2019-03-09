package chapter1arraysstrings

// URLify returns a rune slice with all spaces replaced with %20
//
// Time Complexity: O(n)
// Space Complexity: O(1)
func URLify(s []rune, baseLength int) {
	numSpaces := 0
	for i := 0; i < baseLength; i++ {
		if s[i] == ' ' {
			numSpaces++
		}
	}

	index := baseLength + numSpaces*2
	for i := baseLength - 1; i >= 0; i-- {
		if s[i] == ' ' {
			s[index-1] = '0'
			s[index-2] = '2'
			s[index-3] = '%'
			index = index - 3
		} else {
			s[index-1] = s[i]
			index--
		}
	}
}
