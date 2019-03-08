package chapter1arraysstrings

import (
	"fmt"
	"testing"
)

var tests_1_2 = []struct {
	s1       string
	s2       string
	expected bool
}{
	{"hi", "ih", true},
	{"hello", "bola", false},
}

func TestIsPermutation(t *testing.T) {
	for _, tt := range tests_1_2 {
		testCall := fmt.Sprintf("IsPermutation(%q,%q)", tt.s1, tt.s2)
		t.Run(testCall, func(t *testing.T) {
			res := IsPermutation(tt.s1, tt.s2)
			if tt.expected != res {
				t.Errorf("expected %s == %t, received %t",
					testCall, tt.expected, res)
			}
		})
	}
}
