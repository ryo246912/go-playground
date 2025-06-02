package main

import (
	"encoding/json"
	"fmt"
)

var content = `
{
	"species": "ハト",
	"description": "ハトは、鳥類の一種で、平和の象徴とされています。",
	"dimensions": {
		"height": 40,
		"width": 10
	}
}
`

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
	var data Data
	if err := json.Unmarshal([]byte(content), &data); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Species: %s\n", data.Species)
	fmt.Printf("Description: %s\n", data.Description)
	fmt.Printf("Height: %d, Width: %d\n", data.Dimensions.Height, data.Dimensions.Width)
}
