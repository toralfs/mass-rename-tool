package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println(`
--------------------------------------------
Welcome to Mass Rename Tool
		
Program will look from root directory after 
a specific filename and rename to chosen name
--------------------------------------------
`)
	fmt.Println("Enter file name to look for: ")

	var oldName, newName string
	fmt.Scan(&oldName)
	fmt.Println("Enter new file name: ")
	fmt.Scan(&newName)

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	files, err := walkDir(wd)
	if err != nil {
		log.Fatalln(err)
	}
	renameFiles(files, string(oldName), string(newName))
}

func walkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func renameFiles(files []string, oldName string, newName string) {
	for _, oldPath := range files {
		newPath := strings.ReplaceAll(oldPath, oldName, newName)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			log.Printf("Could not rename file at path: %s, error: %v", oldPath, err)
			continue
		}
	}
}
