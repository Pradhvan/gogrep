package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
		searchWord = strings.ToLower(searchWord)
	}

	var search_line string
	for scanner.Scan() {
		line := scanner.Text()
		if !isCaseSensitive {
			search_line = strings.ToLower(line)
		} else {
			search_line = line
		}
		// Check for case sensivity here
		if strings.Contains(search_line, searchWord) {
			fmt.Println(filepath, line)
		}
	}
}
