package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func findSearchWord(filepath string, searchWord string, isCaseSensitive bool) {
	// Add a check to if we have read permission of a file.
	//fmt.Println(fileInfo.Mode().Perm())
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !isCaseSensitive {
		// (?i) at the beginning of the pattern to make it case
		//insensitive in regex.
		searchWord = "(?i)" + searchWord
	}
	re := regexp.MustCompile(searchWord)

	for scanner.Scan() {
		line := scanner.Text()
		if re.MatchString(line) {
			fmt.Println(filepath, line)
		}
	}
}
