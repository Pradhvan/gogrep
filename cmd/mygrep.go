package cmd

import (
	"fmt"
	"log"
	"regexp"

	"github.com/Pradhvan/gogrep/io"
)

func FindSearchWord(filepath string, searchWord string, isCaseSensitive bool) (matchFound []string, err error) {
	// Add a check to if we have read permission of a file.
	//fmt.Println(fileInfo.Mode().Perm())
	io.CheckFileExists(filepath)
	io.IsDirectory(filepath)
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
	var matchText []string
	var match string
	for _, line := range fileContent {
		if re.MatchString(line) {
			match = fmt.Sprintf("%s: %s", filepath, line)
			matchText = append(matchText, match)
		}
	}
	return matchText, nil
}
