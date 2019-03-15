package chapter1arraysstrings

import (
	"fmt"
	"testing"
)

var tests_1_5 = []struct {
	first    string
	second   string
	expected bool
}{
	{"pale", "ple", true},
	{"pales", "pale", true},
	{"pale", "bale", true},
	{"pale", "bae", false},
}

func TestIsOneAway(t *testing.T) {
	for _, tt := range tests_1_5 {
		testCall := fmt.Sprintf("IsOneAway(%q, %q)", tt.first, tt.second)
		t.Run(testCall, func(t *testing.T) {
			res := IsOneAway(tt.first, tt.second)
			if tt.expected != res {
				t.Errorf("expected %s == %t, received %t",
					testCall, tt.expected, res)
			}
		})
	}
}
