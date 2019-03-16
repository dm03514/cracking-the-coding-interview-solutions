package chapter1arraysstrings

import (
	"strconv"
	"strings"
)

// CompressCharCounts compresses a string collapsing each character using
// its runlength, ie aaaaaa will become a6
//
// TimeComplexity O(n)
func CompressCharCounts(s string) (string, error) {
	var compressed strings.Builder
	consecutiveCount := 0
	for i := 0; i < len(s); i++ {
		consecutiveCount++
		// if the end of the string has been reached OR the next char is different
		if i+1 >= len(s) || s[i] != s[i+1] {
			if err := compressed.WriteByte(s[i]); err != nil {
				return "", err
			}
			if _, err := compressed.WriteString(strconv.Itoa(consecutiveCount)); err != nil {
				return "", err
			}
			consecutiveCount = 0
		}
	}
	c := compressed.String()
	if len(c) < len(s) {
		return c, nil
	}

	return s, nil
}
