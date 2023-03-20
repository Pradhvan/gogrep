package parseflag

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFlagsCorrect(t *testing.T) {
	var tests = []struct {
		name string
		args []string
		conf Config
	}{
		{
			name: "Test when only searchword and filename is passed.",
			args: []string{"foo", "words.txt"},
			conf: Config{
				OutputFile:        "",
				IsCaseSensitive:   false,
				CountSearchResult: false,
				CountBefore:       0,
				Args:              []string{"foo", "words.txt"},
			},
		},
		{
			name: "Test for output flag '-o'",
			args: []string{"-o", "output.txt", "foo", "words.txt"},
			conf: Config{
				OutputFile:        "output.txt",
				IsCaseSensitive:   false,
				CountSearchResult: false,
				CountBefore:       0,
				Args:              []string{"foo", "words.txt"},
			},
		},
		{
			name: "Test for flag '-i'",
			args: []string{"-i", "foo", "words.txt"},
			conf: Config{
				OutputFile:        "",
				IsCaseSensitive:   true,
				CountSearchResult: false,
				CountBefore:       0,
				Args:              []string{"foo", "words.txt"},
			},
		},
		{
			name: "Test for flag '-B'",
			args: []string{"-B", "10", "foo", "words.txt"},
			conf: Config{
				OutputFile:        "",
				IsCaseSensitive:   false,
				CountSearchResult: false,
				CountBefore:       10,
				Args:              []string{"foo", "words.txt"},
			},
		},
		{
			name: "Test for flag '-c'",
			args: []string{"-c", "foo", "words.txt"},
			conf: Config{
				OutputFile:        "",
				IsCaseSensitive:   false,
				CountSearchResult: true,
				CountBefore:       0,
				Args:              []string{"foo", "words.txt"},
			},
		},
		{
			name: "Test to check combined flag usage.",
			args: []string{"-o", "output.txt", "-i", "-c", "-B", "3", "foo", "words.txt"},
			conf: Config{
				OutputFile:        "output.txt",
				IsCaseSensitive:   true,
				CountSearchResult: true,
				CountBefore:       3,
				Args:              []string{"foo", "words.txt"},
			},
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			conf, output, err := ParseFlags("prog", test.args)
			assert.Nil(t, err)
			assert.Equal(t, output, "")
			assert.Equal(t, test.conf, *conf)
		})
	}
}

func TestParseFlagsError(t *testing.T) {
	var tests = []struct {
		name     string
		args     []string
		errorstr string
	}{
		{
			name:     "Test to check error for invalid flag.",
			args:     []string{"-l", "test", "words.txt"},
			errorstr: "flag provided but not defined",
		},
		{
			name:     "Test to check error for invalid input to a valid flag.",
			args:     []string{"-B", "test"},
			errorstr: "invalid value",
		},
		{
			name:     "Test to check error on input.",
			args:     []string{""},
			errorstr: "not enough arguments passed",
		},
		{
			name:     "Test to check error for missing argument searchword and filename.",
			args:     []string{"-i", "-c"},
			errorstr: "missing argument searchword and filename",
		},
		{
			name:     "Test to check error for missing argument filename.",
			args:     []string{"foobar"},
			errorstr: "not enough arguments passed",
		},
	}

	for _, test := range tests {
		t.Run(strings.Join(test.args, " "), func(t *testing.T) {
			conf, _, err := ParseFlags("prog", test.args)
			assert.Nil(t, conf)
			assert.Contains(t, err.Error(), test.errorstr)
		})
	}
}
