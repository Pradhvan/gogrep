package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func findSearchWord(filepath string, searchWord string) {
	// Add a check to if we have read permission of a file.
	//fmt.Println(fileInfo.Mode().Perm())
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	searchWord = strings.ToLower(searchWord)
	var search_line string
	for scanner.Scan() {
		line := scanner.Text()
		search_line = strings.ToLower(line)
		// Check for case sensivity here
		if strings.Contains(search_line, searchWord) {
			fmt.Println(filepath, line)
		}
	}
}
