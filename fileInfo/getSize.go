package fileInfo

import (
	"log"
	"os"
)
//returns Size of a file
func GetSize(pathToFile string) int64 {
	file, err := os.Stat(pathToFile)

	if os.IsNotExist(err) {
		log.Fatal("File does not exsist at: ", pathToFile)
	}

	if err != nil {
		log.Fatal(err)
	}

	return file.Size()
}
