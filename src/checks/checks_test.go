package checks_test

import (
	"github.com/ythosa/disguise/src/checks"
	"testing"
)

func TestCheckFilePrefix(t *testing.T) {
	testCases := []struct {
		input string
		want  error
	}{
		{
			input: "* ",
			want:  nil,
		},
		{
			input: "    ",
			want:  nil,
		},
		{
			input: "",
			want:  checks.InvalidInputError{Name: "file-prefix"},
		},
	}

	for _, tc := range testCases {
		if got := checks.CheckFilePrefix(tc.input); got != tc.want {
			t.Errorf("CheckFilePrefix(%q) = %v", tc.input, got)
		}
	}
}
