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
	conf, output, err := parseFlags(os.Args[0], os.Args[1:])

	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(2)
	} else if err != nil {
		fmt.Println("got error:", err)
		fmt.Println("output: \n", output)
		os.Exit(1)
	}

	if len(conf.args) == 0 {
		fmt.Printf("Usage of our Program: \n")
		fmt.Printf("$ ./mygrep searchword filename.txt  \n")
		os.Exit(1)
	}

	searchWord := conf.args[0]
	fileToSearch := conf.args[0]
	var searchList = []string{}

	directory, err := io.IsDirectory(fileToSearch)
	if err != nil {
		log.Fatal(err)
	}
	if directory {
		searchList, err = io.ListFilesInDir(fileToSearch)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		searchList = append(searchList, fileToSearch)
	}
	if conf.outputFile != "" {
		outFileExists, err := io.CheckFileExists(conf.outputFile)

		if outFileExists {
			log.Fatalf("Error: %s already exists in the current directory.", conf.outputFile)
		}

		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				log.Fatal(err)
			}
		}
	}

	var result = []string{}
	for _, file := range searchList {
		matchFound, _ := cmd.FindSearchWord(file, searchWord, conf.isCaseSensitive, conf.countBefore)
		result = append(result, matchFound...)
	}

	if conf.outputFile == "" && !conf.countSearchResult {
		for _, line := range result {
			fmt.Println(line)
		}
	}
	if conf.outputFile != "" {
		io.WriteToFile(conf.outputFile, result)
	}
	if conf.countSearchResult {
		fmt.Printf("Total matches found for '%s' in %s: %d \n", searchWord, fileToSearch, len(result))
	}
}
