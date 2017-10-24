package fileInfo

import (
	"crypto/sha1"
	"io"
	"log"
	"os"
)

func GetSHA1(pathToFile string) []byte {
	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}