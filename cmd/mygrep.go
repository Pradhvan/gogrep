package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/Pradhvan/gogrep/pkg/ds"
	"github.com/Pradhvan/gogrep/pkg/io"
	"github.com/Pradhvan/gogrep/pkg/parseflag"
)

type Result struct {
	MatchText      []string
	Count          int
	ShowCount      bool
	MatchFileWrote bool
}

func FindSearchWord(config parseflag.Config) (result *Result, err error) {
	var matchStore Result
	searchWord := config.Args[0]
	filepath := config.Args[1]

	exsits, err := io.CheckFileExists(filepath)
	if !exsits {
		return nil, fmt.Errorf("error: %s does not exsists", filepath)
	} else if err != nil {
		return nil, err
	}

	var searchList = []string{}
	isDir, err := io.IsDirectory(filepath)
	if err != nil {
		return nil, err
	}

	if isDir {
		searchList, err = io.ListFilesInDir(filepath)
		if err != nil {
			return nil, err
		}
	} else {
		searchList = append(searchList, filepath)
	}

	if !config.IsCaseSensitive {
		// (?i) at the beginning of the pattern to make it case
		//insensitive in regex.
		searchWord = "(?i)" + searchWord
	}
	re := regexp.MustCompile(searchWord)

	var matchText = []string{}
	var beforeStorage = ds.Queue{}

	var shouldCountBefore = false
	if config.CountBefore != 0 {
		shouldCountBefore = true
	}

	for _, searchfile := range searchList {
		file, err := os.Open(searchfile)
		if err != nil {
			if os.IsPermission(err) {
				return nil, fmt.Errorf("read permission denied for %s", filepath)
			}
			return nil, err
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)

		for scanner.Scan() {
			line := scanner.Text()
			if re.MatchString(line) {
				if len(beforeStorage.GetAll()) > 0 && shouldCountBefore {
					matchText = append(matchText, beforeStorage.GetAll()...)
					beforeStorage.Clear()
				}
				matchText = append(matchText, fmt.Sprintf("%s: %s", filepath, line))
			} else if shouldCountBefore {

				if len(beforeStorage.GetAll()) < config.CountBefore {
					beforeStorage.Enqueue(fmt.Sprintf("%s: %s", filepath, line))
				} else {
					beforeStorage.Dequeue()
					beforeStorage.Enqueue(fmt.Sprintf("%s: %s", filepath, line))
				}
			}
		}
		beforeStorage.Clear()
	}
	matchStore.MatchText = matchText

	if config.OutputFile != "" {
		outFileExists, err := io.CheckFileExists(config.OutputFile)

		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				return nil, err
			}
		}

		if outFileExists {
			return nil, fmt.Errorf("error: %s already exists in the current directory", config.OutputFile)
		}

		err = io.WriteToFile(config.OutputFile, matchText)
		if err != nil {
			return nil, err
		}
		matchStore.MatchFileWrote = true
	}

	if config.CountSearchResult {
		matchStore.Count = len(matchText)
		matchStore.ShowCount = true
	}

	return &matchStore, nil
}
