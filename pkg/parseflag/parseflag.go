package parseflag

import (
	"bytes"
	"flag"
	"fmt"
)

type Config struct {
	OutputFile        string
	IsCaseSensitive   bool
	CountSearchResult bool
	CountBefore       int

	// args are the positional (non-flag) command-line arguments.
	Args []string
}

func ParseFlags(searchWord string, args []string) (config *Config, output string, err error) {

	if len(args) < 2 {
		return nil, "", fmt.Errorf("not enough arguments passed, args: %q", args)
	}

	flags := flag.NewFlagSet(searchWord, flag.ContinueOnError)
	var buf bytes.Buffer
	flags.SetOutput(&buf)

	var conf Config
	flags.BoolVar(&conf.IsCaseSensitive, "i", false, "Make the seach case sensitive.")
	flags.BoolVar(&conf.CountSearchResult, "c", false, "Count number of matches.")
	flags.StringVar(&conf.OutputFile, "o", "", "Filename to store the search results.")
	flags.IntVar(&conf.CountBefore, "B", 0, "Print 'n' lines before the match")

	err = flags.Parse(args)
	if err != nil {
		return nil, buf.String(), err
	}
	if len(flags.Args()) == 0 {
		return nil, "", fmt.Errorf("missing argument searchword and filename")
	}

	conf.Args = flags.Args()
	return &conf, buf.String(), nil
}
