package chapter1arraysstrings

import "testing"

var tests_1_1 = []struct {
	input    string
	expected bool
}{
	{"hi", true},
	{"hello", false},
}

func TestIsUnique(t *testing.T) {
	for _, tt := range tests_1_1 {
		t.Run(tt.input, func(t *testing.T) {
			res := IsUnique(tt.input)
			if tt.expected != res {
				t.Errorf("expected IsUnique(%q) == %t, received %t",
					tt.input, tt.expected, res)
			}
		})
	}
}
