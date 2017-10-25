package fileInfo

import (
	"crypto/md5"
	"io"
	"log"
	"os"
)

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
