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
		{
			name:            "Test for `-B` count before flag with one match",
			path:            "testdata/data.txt",
			searchWord:      "Another",
			isCaseSensitive: true,
			output:          []string{"testdata/data.txt: this is a text here.", "testdata/data.txt: Another text goes here."},
			countBefore:     1,
		},
		{
			name:            "Test for `-B` count before flag with multtiple matchs",
			path:            "testdata/data.txt",
			searchWord:      "program",
			isCaseSensitive: false,
			output: []string{
				"testdata/data.txt: An Idomatic Approach to",
				"testdata/data.txt: Real-World Go Programming",
				"testdata/data.txt: ",
				"testdata/data.txt: Welcome to a tour of the Go programming language.",
				"testdata/data.txt: next or PageDown to go to the next page.",
				"testdata/data.txt: The tour is interactive. Click the Run button now (or press Shift + Enter) to compile and run the program on a remote server. The result is displayed below the code.",
				"testdata/data.txt: ",
				"testdata/data.txt: These example programs demonstrate different aspects of Go. The programs in the tour are meant to be starting points for your own experimentation.",
				"testdata/data.txt: ",
				"testdata/data.txt: Edit the program and run it again.",
			},
			countBefore: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := cmd.FindSearchWord(test.path, test.searchWord, test.isCaseSensitive, test.countBefore)
			assert.Equal(t, test.output, got)

		})
	}

}
