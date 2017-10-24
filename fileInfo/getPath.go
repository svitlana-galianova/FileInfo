package fileInfo

import (
	"log"
	"path/filepath"
)

func GetPath(pathToFile string) string {
	absPath, err := filepath.Abs(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	return absPath
}
