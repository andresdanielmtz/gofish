package api

import (
	"fmt"
	"os"
)

// getFileById retrieves the file path by its name.

func getFileById(nameFile string) (string, error) {
	filePath := "images/" + nameFile
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file %s does not exist", nameFile)
	}
	return filePath, nil
}
