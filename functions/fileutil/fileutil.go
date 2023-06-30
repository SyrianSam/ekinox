package fileutil

import (
	"io/ioutil"
	"log"
)

// reads the content of a file and returns it as a string
func ReadFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
	return string(content), nil
}
