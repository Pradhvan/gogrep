package cmd_test

import (
	"os"
	"testing"

	"github.com/Pradhvan/gogrep/cmd"
	"github.com/Pradhvan/gogrep/pkg/iohandler"
	"github.com/Pradhvan/gogrep/pkg/parseflag"
	"github.com/stretchr/testify/assert"
)

func TestFindSearchWord(t *testing.T) {
	tests := []struct {
		name   string
		output []string
		config parseflag.Config
		count  int
	}{
		{
			name:   "Test for search present in the file.",
			output: []string{"testdata/data.txt: Jhon Bodner's Go Book"},
			config: parseflag.Config{
				OutputFile:        "",
				IsCaseSensitive:   false,
				CountSearchResult: false,
				CountBefore:       0,
				Args:              []string{"Jhon", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for no search word present in the file.",
			output: []string{},
			config: parseflag.Config{
				IsCaseSensitive:   false,
				CountBefore:       0,
				OutputFile:        "",
				CountSearchResult: false,
				Args:              []string{"Python", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for multiple search present in the file.",
			output: []string{"testdata/data.txt: foobar", "testdata/data.txt: FOO"},
			config: parseflag.Config{
				IsCaseSensitive:   false,
				CountBefore:       0,
				OutputFile:        "",
				CountSearchResult: false,
				Args:              []string{"foo", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for `-i` casesensitive search.",
			output: []string{"testdata/data.txt: FOO"},
			config: parseflag.Config{
				IsCaseSensitive:   true,
				CountBefore:       0,
				OutputFile:        "",
				CountSearchResult: false,
				Args:              []string{"FOO", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for `-c` with casesensitive search.",
			output: []string{"testdata/data.txt: FOO"},
			count:  1,
			config: parseflag.Config{
				IsCaseSensitive:   true,
				CountBefore:       0,
				OutputFile:        "",
				CountSearchResult: true,
				Args:              []string{"FOO", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for `-c` without casesensitive search.",
			count:  2,
			output: []string{"testdata/data.txt: foobar", "testdata/data.txt: FOO"},
			config: parseflag.Config{
				IsCaseSensitive:   false,
				CountBefore:       0,
				OutputFile:        "",
				CountSearchResult: true,
				Args:              []string{"foo", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for `-o`.",
			output: []string{"testdata/data.txt: foobar", "testdata/data.txt: FOO"},
			config: parseflag.Config{
				IsCaseSensitive:   false,
				CountBefore:       0,
				OutputFile:        "output.txt",
				CountSearchResult: false,
				Args:              []string{"foo", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for `-B` count before flag with one match",
			output: []string{"testdata/data.txt: this is a text here.", "testdata/data.txt: Another text goes here."},
			config: parseflag.Config{
				IsCaseSensitive:   true,
				CountBefore:       1,
				CountSearchResult: false,
				OutputFile:        "",
				Args:              []string{"Another", "testdata/data.txt"},
			},
		},
		{
			name: "Test for `-B` count before flag with multtiple matchs",
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
			config: parseflag.Config{
				IsCaseSensitive:   false,
				CountBefore:       1,
				CountSearchResult: false,
				OutputFile:        "",
				Args:              []string{"program", "testdata/data.txt"},
			},
		},
		{
			name:   "Test for multiple flags usage.",
			output: []string{"testdata/data.txt: this is a text here.", "testdata/data.txt: Another text goes here."},
			count:  2,
			config: parseflag.Config{
				IsCaseSensitive:   true,
				CountBefore:       1,
				CountSearchResult: true,
				OutputFile:        "output.txt",
				Args:              []string{"Another", "testdata/data.txt"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			r, _ := cmd.FindSearchWord(test.config)
			assert.Equal(t, test.output, r.MatchText)

			if test.config.CountSearchResult {
				assert.Equal(t, test.count, r.Count)
			}
			// Check if output file exists
			// and clean it up after the test is run.
			if len(test.config.OutputFile) > 0 {
				assert.Equal(t, true, r.MatchFileWrote)
				outFileExists, err := iohandler.CheckFileExists(test.config.OutputFile)
				assert.Equal(t, nil, err)
				assert.Equal(t, true, outFileExists)
				if outFileExists {
					err := os.Remove(test.config.OutputFile)
					assert.Equal(t, nil, err)
				}
			}
		})
	}

}
