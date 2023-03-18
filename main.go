package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Pradhvan/gogrep/cmd"
	"github.com/Pradhvan/gogrep/pkg/io"
)

func main() {
	var outputFile string
	var isCaseSensitive bool
	var countSearchResult bool
	var countBefore int

	flag.BoolVar(&isCaseSensitive, "i", false, "Make the seach case sensitive.")
	flag.BoolVar(&countSearchResult, "c", false, "Count number of matches.")
	flag.StringVar(&outputFile, "o", "", "Filename to store the search results.")
	flag.IntVar(&countBefore, "B", 0, "Print 'n' lines before the match")

	flag.Parse()
	searchWord := flag.Arg(0)
	fileToSearch := flag.Arg(1)

	if len(searchWord) == 0 || len(fileToSearch) == 0 {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("$ ./mygrep searchword filename.txt")
		return
	}

	if outputFile != "" {
		fmt.Println(outputFile)
		outFileExists, err := io.CheckFileExists(outputFile)

		if outFileExists {
			log.Fatalf("Error: %s already exists in the current directory.", outputFile)
		}

		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				log.Fatal(err)
			}
		}
	}
	result, _ := cmd.FindSearchWord(fileToSearch, searchWord, isCaseSensitive, countBefore)
	if outputFile == "" && !countSearchResult {
		for _, line := range result {
			fmt.Println(line)
		}
	}
	if outputFile != "" {
		io.WriteToFile(outputFile, result)
	}
	if countSearchResult {
		fmt.Printf("Total matches found for '%s' in %s: %d \n", searchWord, fileToSearch, len(result))
	}
}
