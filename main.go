package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	files, err := walkDir("files/")
	if err != nil {
		log.Fatalln(err)
	}
	renameFiles(files, "README.md")
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

func renameFiles(files []string, newName string) {
	for _, v := range files {
		fmt.Println(v)
		xs := strings.SplitAfter(v, `\`)
		if len(xs) > 1 {
			//newPath := fmt.Sprintf("%s%s", strings.Trim(fmt.Sprint(xs[:len(xs)-1]), "[]"), newName)
			newPath := strings.Trim(strings.Join(xs, " "), " ")
			fmt.Println(newPath)

		}
	}
}
