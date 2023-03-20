package cmd

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/Pradhvan/gogrep/pkg/ds"
	"github.com/Pradhvan/gogrep/pkg/io"
)

func FindSearchWord(filepath string, searchWord string, isCaseSensitive bool, countBefore int) (matchFound []string, err error) {

	exsits, err := io.CheckFileExists(filepath)
	if !exsits {
		return nil, fmt.Errorf("error: %s does not exsists", filepath)
	} else if err != nil {
		return nil, err
	}

	isDir, err := io.IsDirectory(filepath)
	if isDir {
		return nil, fmt.Errorf("current file is a directory ")
	} else if err != nil {
		return nil, err
	}

	if !isCaseSensitive {
		// (?i) at the beginning of the pattern to make it case
		//insensitive in regex.
		searchWord = "(?i)" + searchWord
	}
	re := regexp.MustCompile(searchWord)

	file, err := os.Open(filepath)
	if err != nil {
		if os.IsPermission(err) {
			return nil, fmt.Errorf("read permission denied for %s", filepath)
		}
		return nil, err
	}
	defer file.Close()

	var matchText = []string{}
	var beforeStorage = ds.Queue{}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			if len(beforeStorage.GetAll()) > 0 && countBefore != 0 {
				matchText = append(matchText, beforeStorage.GetAll()...)
				beforeStorage.Clear()
			}
			matchText = append(matchText, fmt.Sprintf("%s: %s", filepath, line))
		} else if countBefore != 0 {

			if len(beforeStorage.GetAll()) < countBefore {
				beforeStorage.Enqueue(fmt.Sprintf("%s: %s", filepath, line))
			} else {
				beforeStorage.Dequeue()
				beforeStorage.Enqueue(fmt.Sprintf("%s: %s", filepath, line))
			}
		}
	}

	return matchText, nil
}
