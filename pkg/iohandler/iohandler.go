package iohandler

import (
	"errors"
	"io/fs"
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

func WriteToFile(outputFile string, result []string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	// Deferring unsafe method "Close" on type "*os.File"
	// Deferring a function call ignores its return value, and the Close() method
	//can return errors. For writable files, we should avoid the defer idiom or very infrequent,
	// maddening bugs will occur.
	for _, line := range result {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			file.Close()
			return err
		}
	}
	return file.Close()
}

func ListFilesInDir(root string) ([]string, error) {
	var filePaths []string
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
