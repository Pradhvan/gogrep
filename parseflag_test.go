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
			assert.Nil(t, err)
			assert.Equal(t, output, "")
			assert.Equal(t, test.conf, *conf)
		})
	}
}

func TestParseFlagsError(t *testing.T) {
	var tests = []struct {
		args   []string
		errstr string
	}{
		{[]string{"-l", "test", "words.txt"}, "flag provided but not defined"},
		{[]string{"-B", "test"}, "invalid value"},
		{[]string{""}, "not enough arguments passed"},
		{[]string{"-i", "-c"}, "missing argument searchword and filename"},
		{[]string{"foobar"}, "not enough arguments passed"},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.args, " "), func(t *testing.T) {
			conf, _, err := parseFlags("prog", tt.args)
			assert.Nil(t, conf)
			assert.Contains(t, err.Error(), tt.errstr)
		})
	}
}
