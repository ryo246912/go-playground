package main

import (
	"log"
	"os"
)

func main() {
	filename := "log.txt"

	f, err := os.OpenFile(filename, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("app started")
}
