package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files := []string{}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(cwd, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		fmt.Printf("%d: %s\n", i, f)
	}
}
