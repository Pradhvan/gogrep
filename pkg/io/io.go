package io

import (
	"bufio"
	"errors"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func CheckFileExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return true, nil
}

func IsDirectory(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}
	if fileInfo.IsDir() {
		return true, nil
	}
	return false, nil
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
			log.Fatalf("Error: Read permission denied for %s", filePath)
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

func ListFilesInDir(root string) ([]string, error) {
	filePaths := []string{}
	// filepath.Walk is less efficient than WalkDir, introduced in Go 1.16,
	//which avoids calling os.Lstat on every visited file or directory.
	err := filepath.WalkDir(root, func(path string, fi fs.DirEntry, err error) error {
		if !fi.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}
