package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetConfig(path string) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	dimensions := make([]Dimension, 5, 5)
	err := decoder.Decode(&dimensions)
	if err != nil {
		fmt.Printf("No Errors on Config!\n")
	}
}
