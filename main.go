package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Pradhvan/gogrep/cmd"
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

	match, err := cmd.FindSearchWord(*conf)
	if err != nil {
		log.Fatal(err)
	}

	if match.ShowCount {
		fmt.Printf("Total matches found for '%s' in %s: %d \n", conf.Args[0], conf.Args[1], match.Count)
		return
	}

	if !match.MatchFileWrote {
		for _, line := range match.MatchText {
			fmt.Println(line)
		}
	}
}
