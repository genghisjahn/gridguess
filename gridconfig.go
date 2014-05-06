import (
    "encoding/json"
    "os"
    "fmt"
)

func GetConfig(path string) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	dimensions = make([]Dimension, 5, 5)
}