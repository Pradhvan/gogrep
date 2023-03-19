package cmd_test

import (
	"testing"

	"github.com/Pradhvan/gogrep/cmd"
	"github.com/stretchr/testify/assert"
)

func TestFindSearchWord(t *testing.T) {
	tests := []struct {
		name            string
		path            string
		searchWord      string
		isCaseSensitive bool
		output          []string
		countBefore     int
	}{
		{
			name:            "Test for search present in the file.",
			path:            "testdata/data.txt",
			searchWord:      "Jhon",
			isCaseSensitive: false,
			output:          []string{"testdata/data.txt: Jhon Bodner's Go Book"},
			countBefore:     0,
		},
		{
			name:            "Test for no search word present in the file.",
			path:            "testdata/data.txt",
			searchWord:      "Python",
			isCaseSensitive: false,
			output:          []string{},
			countBefore:     0,
		},
		{
			name:            "Test for multiple search present in the file.",
			path:            "testdata/data.txt",
			searchWord:      "foo",
			isCaseSensitive: false,
			output:          []string{"testdata/data.txt: foobar", "testdata/data.txt: FOO"},
			countBefore:     0,
		},
		{
			name:            "Test for casesensitive search.",
			path:            "testdata/data.txt",
			searchWord:      "FOO",
			isCaseSensitive: true,
			output:          []string{"testdata/data.txt: FOO"},
			countBefore:     0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := cmd.FindSearchWord(test.path, test.searchWord, test.isCaseSensitive, test.countBefore)
			assert.Equal(t, test.output, got)

		})
	}

}
