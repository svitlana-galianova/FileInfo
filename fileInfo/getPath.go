package fileInfo

import (
	"log"
	"path/filepath"
)

//returns path
func GetPath(pathToFile string) string {
	absPath, err := filepath.Abs(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	return absPath
}
