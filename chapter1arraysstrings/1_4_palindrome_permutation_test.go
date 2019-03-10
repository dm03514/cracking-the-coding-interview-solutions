package chapter1arraysstrings

import (
	"fmt"
	"testing"
)

var tests_1_4 = []struct {
	input    string
	expected bool
}{
	{"Tact Coa", true},
	{"racecar", true},
	{"nope", false},
}

func TestIsPalindromePermutation(t *testing.T) {
	for _, tt := range tests_1_4 {
		testCall := fmt.Sprintf("IsPalindromePermutation(%q)", tt.input)
		t.Run(testCall, func(t *testing.T) {
			res := IsPalindromePermutation(tt.input)
			if tt.expected != res {
				t.Errorf("expected %s == %t, received %t",
					testCall, tt.expected, res)
			}
		})
	}
}
