package fileInfo

import (
	"crypto/sha1"
	"io"
	"log"
	"os"
)

//returns SHA1
func GetSHA1(pathToFile string) []byte {
	//open file
	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// generate sha1
	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}