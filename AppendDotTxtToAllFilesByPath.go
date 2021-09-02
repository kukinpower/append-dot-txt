package main

import (
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) == 2 {
		println("dir: " + os.Args[1])

		var path = os.Args[1]
		var files []string

		err, files := getFilesInPath(path, files)

		if err != nil {
			panic(err)
		}

		if appendDotTxt(files) {
			println("Error occurred")
			return
		}
	} else {
		println("Provide only one argument: absolute path to a dir")
		return
	}
}

func getFilesInPath(path string, files []string) (error, []string) {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return err, files
}

func appendDotTxt(files []string) bool {
	for _, file := range files {
		err := os.Rename(file, file+".txt")
		if err != nil {
			return true
		}
	}
	return false
}
