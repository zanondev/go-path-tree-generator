package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func listFiles(path string, prefix string, output *os.File) {
	dirEntries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			fmt.Fprintln(output, prefix+dirEntry.Name())
			newPath := filepath.Join(path, dirEntry.Name())
			listFiles(newPath, prefix+"|---", output)
		} else {
			fmt.Fprintln(output, prefix+dirEntry.Name())
		}
	}
}

func main() {
	// input
	dirPath := "C:\\repos"

	//output
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	listFiles(dirPath, "", outputFile)
}
