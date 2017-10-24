package main
import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"

	// note: package should be named <project-name>/fileInfo
	// if your local project's name is not "FileInfo"
	// please change the name of the package
	"FileInfo/fileInfo"
)
func main() {
	
	//1. get current working directory(wd)
	wd, err := os.Getwd()
	
	if err != nil {
		log.Fatal(err)
	}

	//2. get bytes
	var buffer bytes.Buffer
	buffer.WriteString(wd)
	buffer.WriteString("\\README.md")
	//3. windows file seperator replace
	path := strings.Replace(buffer.String(), "\\", "/", -1)
	
	fmt.Println("File Information")
	fmt.Println("===========================================")
	fmt.Println("Current working directory: "+wd)
	fmt.Println("File path: path-to-file-goes-here")
	fmt.Println("File Name: "+fileInfo.GetName(path))
	fmt.Printf("SHA1: %x", fileInfo.GetSHA1(path))
	
}
