package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	args := os.Args
	searchWord := args[1]
	fmt.Println(searchWord)
	fileToSearch := args[2]
	// File checks can be abstracted out a seprate function
	fileInfo, err := os.Stat(fileToSearch)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File Does not Exsists")
		return
	}
	if fileInfo.IsDir() {
		fmt.Println("File is a directory")
		return
	}
	// Add a check to if we have read permission of a file.
	//fmt.Println(fileInfo.Mode().Perm())
	file, err := os.Open(fileToSearch)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Check for case sensivity here
		if strings.Contains(line, searchWord) {
			fmt.Println("Yes")
			fmt.Println(line)
		}
	}
}
