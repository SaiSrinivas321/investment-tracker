package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// Define the structure for each asset
type Asset struct {
	Tradingsymbol string  `json:"tradingsymbol"`
	Quantity      int     `json:"quantity"`
	AveragePrice  float64 `json:"average_price"`
}

// Function to generate the curl command for each asset
func generateCurlCommand(baseURL string, asset Asset) string {
	investedAmount := float64(asset.Quantity) * asset.AveragePrice
	accountName := "zerodha" // Fixed account name

	curlCommand := fmt.Sprintf(`
curl --location '%s/investments' \
--header 'Content-Type: application/json' \
--data '{
    "asset_type": "stock",
    "asset_name": "%s",
    "quantity": %d,
    "invested_amount": %.2f,
    "account_name": "%s"
}'`, baseURL, asset.Tradingsymbol, asset.Quantity, investedAmount, accountName)

	return curlCommand
}

// Function to execute the curl command
func executeCurlCommand(curlCommand string) error {
	// Split the curl command into arguments for exec.Cmd
	args := []string{"-c", curlCommand}

	// Create the command using exec.Command
	cmd := exec.Command("sh", args...)

	// Run the command and capture any error
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute curl command: %w", err)
	}
	return nil
}

func main() {
	// Open the JSON file
	file, err := os.Open("assets.json") // Change "assets.json" to your file name
	if err != nil {
		log.Fatalf("Error opening JSON file: %v", err)
	}
	defer file.Close()

	// Read the content of the file
	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading file content: %v", err)
	}

	// Parse the JSON data into a slice of Asset structs
	var assets []Asset
	err = json.Unmarshal(fileContent, &assets)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	// Base URL for the API
	baseURL := "<base_url>" // Replace with your actual API URL

	// Loop through the assets and generate and execute the curl command for each
	for _, asset := range assets {
		curlCommand := generateCurlCommand(baseURL, asset)

		// Execute the curl command
		err := executeCurlCommand(curlCommand)
		if err != nil {
			log.Printf("Error executing curl command for asset %s: %v", asset.Tradingsymbol, err)
		} else {
			fmt.Printf("Successfully executed curl for asset: %s\n", asset.Tradingsymbol)
		}
	}
}
