package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/Pradhvan/gogrep/cmd"
	"github.com/Pradhvan/gogrep/io"
)

func main() {
	isCaseSensitive := flag.Bool("i", false, "Make the seach case sensitive.")
	outputFile := flag.String("o", "", "Filename to store the search results.")
	flag.Parse()
	searchWord := flag.Arg(0)
	fileToSearch := flag.Arg(1)
	//File checks can be abstracted out a seprate function
	fileInfo, err := os.Stat(fileToSearch)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File Does not Exsists")
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("File is a directory")
		return
	}
	result, _ := cmd.FindSearchWord(fileToSearch, searchWord, *isCaseSensitive)
	for _, line := range result {
		fmt.Println(line)
	}
	if *outputFile != "" {
		io.WriteToFile(*outputFile, result)
	}
}
