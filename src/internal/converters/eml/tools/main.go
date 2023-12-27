package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/alcionai/corso/src/internal/common/sanitize"
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

	fmt.Printf("Original %s\n", string(byteValue))

	// Declare an empty interface for holding the JSON data
	var jsonData any

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
	}

	sanitizedValue := sanitize.JSONBytes(byteValue)

	fmt.Printf("\nSanitized %s\n", string(sanitizedValue))

	if !json.Valid(sanitizedValue) {
		fmt.Println("sanitizedValue INVALID JSON")
	}

	_, err = api.BytesToMessageable(sanitizedValue)
	if err != nil {
		fmt.Println("sanitizedValue: Error converting to messagable", err)
	}

	// Unmarshal the byteValue into the jsonData interface
	err = json.Unmarshal(sanitizedValue, &jsonData)
	if err != nil {
		fmt.Println("sanitizedValue: Error parsing JSON:", err)
	}
}
