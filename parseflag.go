package main

import (
	"bytes"
	"flag"
	"fmt"
)

type Config struct {
	outputFile        string
	isCaseSensitive   bool
	countSearchResult bool
	countBefore       int

	// args are the positional (non-flag) command-line arguments.
	args []string
}

func parseFlags(searchWord string, args []string) (config *Config, output string, err error) {

	if len(args) < 2 {
		return nil, "", fmt.Errorf("not enough arguments passed, args: %q", args)
	}

	flags := flag.NewFlagSet(searchWord, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf Config
	flags.BoolVar(&conf.isCaseSensitive, "i", false, "Make the seach case sensitive.")
	flags.BoolVar(&conf.countSearchResult, "c", false, "Count number of matches.")
	flags.StringVar(&conf.outputFile, "o", "", "Filename to store the search results.")
	flags.IntVar(&conf.countBefore, "B", 0, "Print 'n' lines before the match")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	if len(flags.Args()) == 0 {
		return nil, "", fmt.Errorf("missing argument searchword and filename")
	}

	conf.args = flags.Args()
	return &conf, buf.String(), nil
}
