package chapter1arraysstrings

import (
	"fmt"
	"testing"
)

var tests_1_3 = []struct {
	input      []rune
	baseLength int
	expected   []rune
}{
	{[]rune("Mr John Smith    "), 13, []rune("Mr%20John%20Smith")},
}

func TestURLify(t *testing.T) {
	for _, tt := range tests_1_3 {
		testCall := fmt.Sprintf("URLify(%q, %d)", tt.input, tt.baseLength)
		t.Run(testCall, func(t *testing.T) {
			URLify(tt.input, tt.baseLength)
			if string(tt.expected) != string(tt.input) {
				t.Errorf("expected %s == %q, received %q",
					testCall, tt.expected, tt.input)
			}
		})
	}
}
