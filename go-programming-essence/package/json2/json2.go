package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Dimensions struct {
	Height int `json:"height"`
	Width  int `json:"width"`
}

type Data struct {
	Species     string     `json:"species"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func main() {
	f, err := os.Open("test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data Data
	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Species: %s\n", data.Species)
	fmt.Printf("Description: %s\n", data.Description)
	fmt.Printf("Height: %d, Width: %d\n", data.Dimensions.Height, data.Dimensions.Width)
}
