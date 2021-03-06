package main

import (
	"bufio"
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
--------------------------------------------`)

	fmt.Println("\n\nEnter file name to change: ")

	oldName := readInput()
	fmt.Println("Enter new file name: ")
	newName := readInput()

	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	// walk through directories from wd and check if matching files were found
	files, err := walkDir(wd, oldName)
	if err != nil {
		log.Fatalln(err)
	}
	if len(files) == 0 {
		fmt.Println("\nNo files found, press anything to exit")
		readInput()
		return
	}

	fmt.Println("\nFound these files: ")
	for _, v := range files {
		fmt.Println(v)
	}
	fmt.Printf("\nSure you want to rename to %s?\nPress y, or n to continue\n", newName)

	ans := readInput()
	if ans == "n" {
		fmt.Println("Renaming cancelled, press anything to exit.")
		readInput()
	} else if ans == "y" {
		errs := renameFiles(files, string(oldName), string(newName))
		if len(errs) > 0 {
			fmt.Println("Some files could not be renamed, check log for details.")
		}
		fmt.Println("All files renamed sucessfully!")
	}
	fmt.Println("press anything to exit")
	readInput()
}

func walkDir(root string, fileName string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && filepath.Base(path) == fileName {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func renameFiles(files []string, oldName string, newName string) []error {
	var errs []error
	for _, oldPath := range files {
		newPath := strings.ReplaceAll(oldPath, oldName, newName)
		err := os.Rename(oldPath, newPath)
		if err != nil {
			e := fmt.Errorf("Could not rename file at path: %s, error: %v", oldPath, err)
			fmt.Println(e)
			errs = append(errs, e)
			continue
		}
	}
	return errs
}

func readInput() string {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	ln := s.Text()
	if err := s.Err(); err != nil {
		log.Println("could not read input: ", err)
	}
	return ln
}
