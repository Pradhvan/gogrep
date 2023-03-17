package io

import (
	"bufio"
	"errors"
	"log"
	"os"
)

func CheckFileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return true, nil
	}
}

func IsDirectory(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.IsDir() {
		return true, nil
	} else {
		return false, nil
	}
}

func WriteToFile(outputFile string, result []string) {
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for _, line := range result {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ReadFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsPermission(err) {
			log.Println("Error: Read permission denied.")
		} else {
			log.Fatal(err)
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var fileContent []string
	for scanner.Scan() {
		fileContent = append(fileContent, scanner.Text())
	}
	return fileContent, scanner.Err()
}
