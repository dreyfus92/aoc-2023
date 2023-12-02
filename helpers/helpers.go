package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetInput(day string, debug bool) (string, error) {
	// Get current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting the directory:", err)
		return "", err
	}
	// Set Path
	var path string
	if !debug {
		path = fmt.Sprintf("./%s/input.txt", day)
	} else {
		path = "input.txt"
	}

	// Read the text file
	pathToInput := filepath.Join(wd, path)
	fileContent, err := os.ReadFile(pathToInput)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	return string(fileContent), nil
}
