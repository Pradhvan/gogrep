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
	}{
		{
			name:            "Test for search present in the file.",
			path:            "testdata/data.txt",
			searchWord:      "Jhon",
			isCaseSensitive: false,
			output:          []string{"testdata/data.txt: Jhon Bodner's Go Book"},
		},
		{
			name:            "Test for no search word present in the file.",
			path:            "testdata/data.txt",
			searchWord:      "Python",
			isCaseSensitive: false,
			output:          []string{},
		},
		{
			name:            "Test for multiple search present in the file.",
			path:            "testdata/data.txt",
			searchWord:      "foo",
			isCaseSensitive: false,
			output:          []string{"testdata/data.txt: foobar", "testdata/data.txt: FOO"},
		},
		{
			name:            "Test for casesensitive search.",
			path:            "testdata/data.txt",
			searchWord:      "FOO",
			isCaseSensitive: true,
			output:          []string{"testdata/data.txt: FOO"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := cmd.FindSearchWord(test.path, test.searchWord, test.isCaseSensitive)
			assert.Equal(t, test.output, got)

		})
	}

}
