package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Pradhvan/gogrep/cmd"
	"github.com/Pradhvan/gogrep/pkg/io"
	"github.com/Pradhvan/gogrep/pkg/parseflag"
)

func main() {
	conf, output, err := parseflag.ParseFlags(os.Args[0], os.Args[1:])

	if err == flag.ErrHelp {
		fmt.Println(output)
		os.Exit(1)
	} else if err != nil {
		fmt.Println("Error: \n", err)
		os.Exit(1)
	}

	searchWord := conf.Args[0]
	fileToSearch := conf.Args[1]
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
	if conf.OutputFile != "" {
		outFileExists, err := io.CheckFileExists(conf.OutputFile)

		if outFileExists {
			log.Fatalf("Error: %s already exists in the current directory.", conf.OutputFile)
		}

		if err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				log.Fatal(err)
			}
		}
	}

	var result = []string{}
	for _, file := range searchList {
		matchFound, _ := cmd.FindSearchWord(file, searchWord, conf.IsCaseSensitive, conf.CountBefore)
		result = append(result, matchFound...)
	}

	if conf.OutputFile == "" && !conf.CountSearchResult {
		for _, line := range result {
			fmt.Println(line)
		}
	}
	if conf.OutputFile != "" {
		io.WriteToFile(conf.OutputFile, result)
	}
	if conf.CountSearchResult {
		fmt.Printf("Total matches found for '%s' in %s: %d \n", searchWord, fileToSearch, len(result))
	}
}
