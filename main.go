package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Pradhvan/gogrep/cmd"
	"github.com/Pradhvan/gogrep/pkg/io"
)

func main() {
	var outputFile string
	var isCaseSensitive bool

	flag.BoolVar(&isCaseSensitive, "i", false, "Make the seach case sensitive.")
	flag.StringVar(&outputFile, "o", "", "Filename to store the search results.")

	flag.Parse()
	searchWord := flag.Arg(0)
	fileToSearch := flag.Arg(1)

	if len(searchWord) == 0 || len(fileToSearch) == 0 {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("$ ./mygrep searchword filename.txt")
		return
	}

	if outputFile != "" {
		outFileExists, err := io.CheckFileExists(outputFile)
		if err != nil {
			log.Fatal(err)
		}

		if outFileExists {
			log.Fatalf("Error: %s already exists in the current directory.", outputFile)
		}
	}
	result, _ := cmd.FindSearchWord(fileToSearch, searchWord, isCaseSensitive)
	if outputFile == "" {
		for _, line := range result {
			fmt.Println(line)
		}
	} else {
		io.WriteToFile(outputFile, result)
	}
}
