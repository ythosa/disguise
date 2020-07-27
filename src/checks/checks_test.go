package checks_test

import (
	"errors"
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
		if got := checks.CheckFilePrefix(tc.input); !errors.Is(got, tc.want) {
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
		if got := checks.CheckFolderPrefix(tc.input); !errors.Is(got, tc.want) {
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
		if got := checks.CheckExtension(tc.input); !errors.Is(got, tc.want) {
			t.Errorf("CheckExtension(%q) = %v", tc.input, got)
		}
	}
}

func TestCheckRepositoryURL(t *testing.T) {
	testCases := []struct {
		input string
		want  error
	}{
		{
			input: "https://github.com/Ythosa/disguise",
			want:  nil,
		},
		{
			input: "https://github.com/Ythosa",
			want:  checks.InvalidInputError{Name: "repository URL"},
		},
		{
			input: "https://golang/something...",
			want:  checks.InvalidInputError{Name: "repository URL"},
		},
		{
			input: "",
			want:  checks.InvalidInputError{Name: "repository URL"},
		},
		{
			input: "https://github.com/disguise/Ythosa",
			want:  nil,
		},
	}

	for _, tc := range testCases {
		if got := checks.CheckRepositoryURL(tc.input); !errors.Is(got, tc.want) {
			t.Errorf("CheckRepositoryURL(%q) = %v", tc.input, got)
		}
	}
}
