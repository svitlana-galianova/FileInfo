package fileInfo

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)
// returns MD5
func GetMD5(pathToFile string) []byte {
	f, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}
