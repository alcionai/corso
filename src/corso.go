package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Replace with your actual access token and file ID
	accessToken := ""
	fileID := "012V4FEN5TABC2KSIDJVD24MVPIYCFGAUC"

	url := fmt.Sprintf("https://graph.microsoft.com/v1.0/drives/b!I8bdNt6HiEy0o9al4-dl3w5Gk8jIUJdEjzqonlSRf48i67LJdwopT4-6kiycJ5AV/items/%s/content", fileID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error:", resp.Status)
		return
	}

	// Create a new file to write the downloaded content
	file, err := os.Create("local_filename.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Use io.Copy to efficiently copy the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error copying content to file:", err)
		return
	}

	fmt.Println("File downloaded successfully.")
}
