package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
)

/*
go:embed provides some functionality to load external files as variable when compile time.
the loaded variable is below the declaration of go:embed.
*/

//go:embed version.txt
var version string // load txt as string

//go:embed files/img/test-img.png
var imgFiles []byte // load test-img as byte

//go:embed files/*.txt
var path embed.FS // using embed file system with wildcard.

func main() {
	fmt.Println(version)

	// try to write loaded byte in test-img to new files
	err := os.WriteFile("files/img/test-img-copy.png", imgFiles, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	// use embed FS as path to load all files in files directory
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
