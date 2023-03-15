package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	isCaseSensitive := flag.Bool("i", false, "Make the seach case sensitive.")
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
	findSearchWord(fileToSearch, searchWord, *isCaseSensitive)
}
