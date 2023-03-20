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
	//os.Args[0] would be ./mygrep
	// so if not value is passed show usage
	if len(os.Args) < 3 {
		fmt.Printf("Usage: \n")
		fmt.Printf("$ ./mygrep -flag[Optional Value] searchword filename.txt \n")
		os.Exit(1)
	}

	conf, output, err := parseFlags(os.Args[0], os.Args[1:])

	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Error: \n", err)
		os.Exit(1)
	} else if len(conf.args) == 0 {
		fmt.Printf("Usage: \n")
		fmt.Printf("$ ./mygrep -flag[Optional Value] searchword filename.txt \n")
		os.Exit(1)
	}

	searchWord := conf.args[0]
	fileToSearch := conf.args[1]
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
