package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	searchWord := args[1]
	fileToSearch := args[2]
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
	findSearchWord(fileToSearch, searchWord)
}
