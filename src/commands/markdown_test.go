package commands_test

import (
	"github.com/ythosa/disguise/src/commands"
	"testing"
)

func TestIsContains(t *testing.T) {
	type tcase struct {
		elements	[]string
		pattern		string
	}

	testCases := []struct{
		input tcase
		want bool
	}{
		{
			input: tcase{
				elements: []string{"/src"},
				pattern: "https://github.com/Ythosa/disguise/src",
			},
			want: true,
		},
	}

	for _, tc := range testCases {
		if got := commands.IsContains(tc.input.elements, tc.input.pattern); got != tc.want {
			t.Errorf("IsContains(%q) = %v", tc.input, got)
		}
	}
}
