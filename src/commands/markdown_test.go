package commands_test

import (
	"testing"

	"github.com/ythosa/disguise/src/commands"
)

func TestIsContains(t *testing.T) {
	type tcase struct {
		elements []string
		some     string
	}

	testCases := []struct {
		input tcase
		want  bool
	}{
		{
			input: tcase{
				elements: []string{"/src"},
				some:     "https://github.com/Ythosa/disguise/src",
			},
			want: true,
		},
		{
			input: tcase{
				elements: []string{"/src/app"},
				some:     "https://github.com/Ythosa/disguise/src",
			},
			want: false,
		},
		{
			input: tcase{
				elements: []string{"tests", "benchmarks"},
				some:     "https://github.com/linksplatform/data.doublets",
			},
			want: false,
		},
		{
			input: tcase{
				elements: []string{"tests", "benchmarks"},
				some:     "https://github.com/linksplatform/data.doublets/data.tests",
			},
			want: true,
		},
		{
			input: tcase{
				elements: []string{},
				some:     "https://github.com/linksplatform/data.doublets/data.tests",
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		if got := commands.IsContains(tc.input.elements, tc.input.some); got != tc.want {
			t.Errorf("IsContains(%q) = %v", tc.input, got)
		}
	}
}

func TestParseHrefAttr(t *testing.T) {
	type inputT struct {
		href      string
		extension string
	}

	type outputT struct {
		isDir         bool
		isTrackedFile bool
		dirname       string
	}

	testCases := []struct {
		input inputT
		want  outputT
	}{
		{
			input: inputT{
				href:      "https://github.com/Ythosa/whynote/tree/master/app",
				extension: ".js",
			},
			want: outputT{
				isDir:         true,
				isTrackedFile: false,
				dirname:       "app",
			},
		},
		{
			input: inputT{
				href:      "https://github.com/Ythosa/where-is/blob/master/src/app.js",
				extension: ".js",
			},
			want: outputT{
				isDir:         false,
				isTrackedFile: true,
				dirname:       "src",
			},
		},
		{
			input: inputT{
				href:      "https://github.com/Ythosa/where-is/tree/master/src/configs",
				extension: ".py",
			},
			want: outputT{
				isDir:         true,
				isTrackedFile: false,
				dirname:       "src/configs",
			},
		},
		{
			input: inputT{
				href:      "https://github.com/Ythosa/where-is/blob/master/src/configs/config.js",
				extension: ".py",
			},
			want: outputT{
				isDir:         false,
				isTrackedFile: false,
				dirname:       "",
			},
		},
		{
			input: inputT{
				href:      "https://github.com/Ythosa/where-is/blob/master/src/configs/config.js",
				extension: ".js",
			},
			want: outputT{
				isDir:         false,
				isTrackedFile: true,
				dirname:       "src/configs",
			},
		}, {
			input: inputT{
				href:      "https://github.com/Ythosa/where-is/blob/master/package.json",
				extension: ".json",
			},
			want: outputT{
				isDir:         false,
				isTrackedFile: true,
				dirname:       "/",
			},
		},
	}

	for _, tc := range testCases {
		if isDir, isTrackedFile, dirname := commands.ParseHrefAttr(tc.input.href, tc.input.extension);
			isDir != tc.want.isDir || isTrackedFile != tc.want.isTrackedFile || dirname != tc.want.dirname {
			t.Errorf("ParseHrefAttr(%q, %q) = %v, %v, %v", tc.input.href, tc.input.extension,
				isDir, isTrackedFile, dirname)
		}
	}
}

func TestGetDirHref(t *testing.T) {
	type inputT struct {
		filehref string
		dirname string
	}

	testCases := []struct{
		input inputT
		want string
	}{
		{
			input: inputT{
				filehref: "https://github.com/Ythosa/where-is/blob/master/src/libs/printer.js",
				dirname: "src/libs",
			},
			want: "https://github.com/Ythosa/where-is/tree/master/src/libs",
		},
		{
			input: inputT{
				filehref: "https://github.com/Ythosa/where-is/blob/master/package.json",
				dirname: "/",
			},
			want: "https://github.com/Ythosa/where-is/tree/master",
		},
	}

	for _, tc := range testCases {
		if got := commands.GetDirHref(tc.input.filehref, tc.input.dirname); got != tc.want {
			t.Errorf("GetDirHref(%q, %q) = %v", tc.input.filehref, tc.input.dirname, got)
		}
	}
}
