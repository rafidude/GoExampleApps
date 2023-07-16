package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	data := map[string]any{}
	file, err := os.ReadFile("data.json")

	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(string(file))

	err = json.Unmarshal(file, &data)
	if err != nil {
		fmt.Println("Error unmarshalling file:", err)
		return
	}

	fmt.Println(data)
}
