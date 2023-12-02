package helpers

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetInput(day string) (string, error) {
	// Get current working directory
	wd, _ := os.Getwd()

	// Set debug path
	path := "input.txt"

	// Read the text file
	pathToInput := filepath.Join(wd, path)
	fileContent, err := os.ReadFile(pathToInput)
	if err != nil {
		fmt.Println("Error reading file:", err)
		fmt.Println("The path is:", pathToInput)
		return "", err
	}

	return string(fileContent), nil
}
