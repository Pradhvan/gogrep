package cmd

import (
	"fmt"
	"log"
	"regexp"

	"github.com/Pradhvan/gogrep/pkg/ds"
	"github.com/Pradhvan/gogrep/pkg/io"
)

func FindSearchWord(filepath string, searchWord string, isCaseSensitive bool, countBefore int) (matchFound []string, err error) {
	exsits, err := io.CheckFileExists(filepath)
	if !exsits {
		log.Fatalf("Error: %s does not exsists.", filepath)
	}
	if err != nil {
		log.Fatal(err)
	}
	isDir, err := io.IsDirectory(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if isDir {
		log.Fatal("Error: Current file is a directory.")
	}
	fileContent, err := io.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if !isCaseSensitive {
		// (?i) at the beginning of the pattern to make it case
		//insensitive in regex.
		searchWord = "(?i)" + searchWord
	}
	re := regexp.MustCompile(searchWord)
	var matchText = []string{}
	var match string
	var beforeStorage = ds.Queue{}
	if countBefore == 0 {
		for _, line := range fileContent {
			if re.MatchString(line) {
				match = fmt.Sprintf("%s: %s", filepath, line)
				matchText = append(matchText, match)
			}
		}
	} else {
		for _, line := range fileContent {
			if re.MatchString(line) {
				if len(beforeStorage.GetAll()) > 0 {
					matchText = append(matchText, beforeStorage.GetAll()...)
					beforeStorage.Clear()
				}
				match = fmt.Sprintf("%s: %s", filepath, line)
				matchText = append(matchText, match)
			} else {
				if len(beforeStorage.GetAll()) < countBefore {
					beforeStorage.Enqueue(line)
				} else {
					beforeStorage.Dequeue()
					beforeStorage.Enqueue(line)
				}
			}
		}
	}

	return matchText, nil
}
