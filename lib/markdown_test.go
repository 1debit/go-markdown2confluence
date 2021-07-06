package lib

import (
    "testing"
	"strings"
	"path/filepath"
	"fmt"
)

func TestGetParentsFromPath(t *testing.T) {
	var tests = []struct {
		input string
		expected []string
	}{
		{"x/y/z.md", []string{"x", "x / y"}},
		{"x.md", []string{}},
		{"./x/y/z.md", []string{"x", "x / y"}},
	}

	for _, test := range tests {
		parents := getParentsFromPath(test.input)
		if parents == nil && test.expected == nil {
			continue
		}
		if len(parents) != len(test.expected) {
			t.Errorf("expected: %d, got %d", len(test.expected), len(parents))
		}
		for idx, fqParentName := range parents {
			if fqParentName != test.expected[idx] {
				t.Errorf("expected: %s, for idx %d got %s", test.expected[idx], idx, fqParentName)
			}
		}
	}
	
}

// adding this to confirm my understanding of old code 
func TestOldParentsCode(t *testing.T) {
	var tests = []struct {
		inputPath string
		filePath string
		expected []string
	}{
		{".", "x/y/z.md", []string{"x", "y"}},
		{".", "x.md", []string{}},
		{"./docs", "./docs/x/y/z.md", []string{"x", "y"}},
	}

	for _, test := range tests {
		f := test.inputPath
		path := test.filePath
		result := deleteFromSlice(strings.Split(filepath.Dir(strings.TrimPrefix(filepath.ToSlash(path), filepath.ToSlash(f))), "/"), ".")
		// do this here because it was handled elsewhere in the code
		result = deleteFromSlice(result, "")
		fmt.Println(path, result)
		if len(result) != len(test.expected) {
			t.Errorf("expected: %d, got %d", len(test.expected), len(result))
		}
		for idx, parentName := range result {
			if parentName != test.expected[idx] {
				t.Errorf("expected: %s, for idx %d got %s", test.expected[idx], idx, parentName)
			}
		}
	}
}