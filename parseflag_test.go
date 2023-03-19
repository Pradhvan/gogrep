package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFlagsCorrect(t *testing.T) {
	var tests = []struct {
		args []string
		conf Config
	}{
		{
			[]string{"foo", "words.txt"},
			Config{
				outputFile:        "",
				isCaseSensitive:   false,
				countSearchResult: false,
				countBefore:       0,
				args:              []string{"foo", "words.txt"},
			},
		},
		{
			[]string{"-o", "output.txt", "foo", "words.txt"},
			Config{
				outputFile:        "output.txt",
				isCaseSensitive:   false,
				countSearchResult: false,
				countBefore:       0,
				args:              []string{"foo", "words.txt"},
			},
		},
		{
			[]string{"-i", "foo", "words.txt"},
			Config{
				outputFile:        "",
				isCaseSensitive:   true,
				countSearchResult: false,
				countBefore:       0,
				args:              []string{"foo", "words.txt"},
			},
		},
		{
			[]string{"-B", "10", "foo", "words.txt"},
			Config{
				outputFile:        "",
				isCaseSensitive:   false,
				countSearchResult: false,
				countBefore:       10,
				args:              []string{"foo", "words.txt"},
			},
		},
		{
			[]string{"-c", "foo", "words.txt"},
			Config{
				outputFile:        "",
				isCaseSensitive:   false,
				countSearchResult: true,
				countBefore:       0,
				args:              []string{"foo", "words.txt"},
			},
		},
		{
			[]string{"-o", "output.txt", "-i", "-c", "-B", "3", "foo", "words.txt"},
			Config{
				outputFile:        "output.txt",
				isCaseSensitive:   true,
				countSearchResult: true,
				countBefore:       3,
				args:              []string{"foo", "words.txt"},
			},
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			conf, output, err := parseFlags("prog", test.args)
			if err != nil {
				t.Errorf("err got %v, want nil", err)
			}
			assert.Equal(t, output, "")

			assert.Equal(t, test.conf, *conf)
		})
	}
}
