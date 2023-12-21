package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

func main() {
	// Open the JSON file
	jsonFile, err := os.Open("data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer jsonFile.Close()

	// Read the file contents
	byteValue, _ := io.ReadAll(jsonFile)

	// Declare an empty interface for holding the JSON data
	var jsonData interface{}

	if !json.Valid(byteValue) {
		fmt.Println("INVALID JSON")
	}

	_, err = api.BytesToMessageable(byteValue)
	if err != nil {
		fmt.Println("Error converting to messagable", err)
	}

	// Unmarshal the byteValue into the jsonData interface
	err = json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Marshal the data back into JSON for pretty output
	prettyJSON, err := json.MarshalIndent(jsonData, "", "    ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	// Print the pretty JSON
	fmt.Println(string(prettyJSON))

	if !json.Valid(byteValue) {
		fmt.Println("INVALID JSON")
	}

	_, err = api.BytesToMessageable(byteValue)
	if err != nil {
		fmt.Println("Error converting to messagable", err)
	}
	fmt.Println()
}
