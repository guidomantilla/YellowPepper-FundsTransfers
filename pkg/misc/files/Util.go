package files

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFile(fullFileName string) *[]string {

	lines := make([]string, 0)

	file, err := os.Open(fullFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &lines
}

func ValidateIfFolderExists(directory string) bool {
	_, err := os.Stat(directory)
	if err != nil {
		return false
	}
	return true
}

