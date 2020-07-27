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

func TestCheckFolderPrefix(t *testing.T) {
	testCases := []struct {
		input string
		want  error
	}{
		{
			input: "***",
			want:  nil,
		},
		{
			input: "#### ",
			want:  nil,
		},
		{
			input: "",
			want:  checks.InvalidInputError{Name: "folder-prefix"},
		},
	}

	for _, tc := range testCases {
		if got := checks.CheckFolderPrefix(tc.input); got != tc.want {
			t.Errorf("CheckFolderPrefix(%q) = %v", tc.input, got)
		}
	}
}

func TestCheckExtension(t *testing.T) {
	testCases := []struct {
		input string
		want  error
	}{
		{
			input: ".go",
			want:  nil,
		},
		{
			input: " ",
			want:  checks.InvalidInputError{Name: "file extension"},
		},
		{
			input: "js",
			want:  checks.InvalidInputError{Name: "file extension"},
		},
	}

	for _, tc := range testCases {
		if got := checks.CheckExtension(tc.input); got != tc.want {
			t.Errorf("CheckExtension(%q) = %v", tc.input, got)
		}
	}
}
