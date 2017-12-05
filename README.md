# FileInfo 
[![Build Status](https://travis-ci.org/svitlana-galianova/FileInfo.svg?branch=master)](https://travis-ci.org/svitlana-galianova/FileInfo.svg?branch=master)


### Summary
Library for Obtaining File Info Written in Go(Golang)    

This library has a set of functions that generate/display the following information:   
- File Path (fileInfo/getPath.go)
- File Name (fileInfo/getName.go)
- File size (fileInfo/getSize.go)
- SHA1 (fileInfo/getSha1.go)
- MD5 (fileInfo/getMd5.go)

### Set up
1. get [golang](https://golang.org/dl/)
2. clone project locally to your `%GOPATH%/src`
3. note: if you renamed the project locally, please rename package name in `main.go` review code comments in **import** section for more information
4. get dependencies: `go get ./...`
5. run application: `go run main.go`

### Notes
- Please add every new function to the `fileInfo` folder using meaningful name
- **Every exported function's name must start with upper case letter**
- All the functions should be tested and produce output in main.go with a meaningful message

### License
[MIT License](https://github.com/svitlana-galianova/FileInfo/blob/master/LICENSE)
