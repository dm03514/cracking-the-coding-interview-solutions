package chapter1arraysstrings

import (
	"fmt"
	"testing"
)

var tests_1_6 = []struct {
	input    string
	expected string
}{
	{"aabcccccaaa", "a2b1c5a3"},
}

func TestCompressCharCounts(t *testing.T) {
	for _, tt := range tests_1_6 {
		testCall := fmt.Sprintf("CompressCharCounts(%q)", tt.input)
		t.Run(testCall, func(t *testing.T) {
			res, _ := CompressCharCounts(tt.input)
			if tt.expected != res {
				t.Errorf("expected %s == %q, received %q",
					testCall, tt.expected, res)
			}
		})
	}
}
