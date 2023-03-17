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

	flag.BoolVar(&isCaseSensitive, "i", false, "Make the seach case sensitive.")
	flag.StringVar(&outputFile, "o", "", "Filename to store the search results.")

	if outputFile != "" {
		outFileExists, err := io.CheckFileExists(outputFile)

		if err != nil {
			log.Fatal(err)
		}

		if outFileExists {
			log.Fatalf("Error: %s already exists in the current directory.", outputFile)
		}
	}

	flag.Parse()
	searchWord := flag.Arg(0)
	fileToSearch := flag.Arg(1)

	fileInfo, err := os.Stat(fileToSearch)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File Does not Exsists")
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("File is a directory")
		return
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
