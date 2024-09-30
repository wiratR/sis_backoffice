package utils

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// EncodeBase64 encodes a string to Base64 format.
func EncodeBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

// DecodeBase64 decodes a Base64 encoded string.
func DecodeBase64(input string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}

// LogError logs an error message.
func LogError(err error) {
	if err != nil {
		log.Println("Error:", err)
	}
}

// Example utility function to check if a string is empty.
func IsEmpty(s string) bool {
	return len(s) == 0
}

// ListBrands retrieves a list of brand folders in the /image/product/ directory.
func ListBrands() ([]string, error) {
	productDir := filepath.Join("image", "product")

	// Read the contents of the product directory
	entries, err := ioutil.ReadDir(productDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read product directory: %w", err)
	}

	var brands []string
	for _, entry := range entries {
		// Check if the entry is a directory
		if entry.IsDir() {
			brands = append(brands, entry.Name())
		}
	}

	return brands, nil
}

// GetImageData retrieves an image from the specified input image path and returns its Base64-encoded data.
func GetImageData(imagePath string) (string, error) {
	// Open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %w", err)
	}
	defer file.Close()

	// Read the file content into a byte slice
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read image data: %w", err)
	}

	// Encode the image data to Base64
	base64Data := base64.StdEncoding.EncodeToString(imageData)

	return base64Data, nil
}

// CountFilesInDirectory counts the number of files in the specified directory and returns their full paths.
func CountFilesInDirectory(dirPath string) (int, []string, error) {
	// Check if the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return 0, nil, fmt.Errorf("directory does not exist: %w", err)
	}

	// Read the contents of the directory
	entries, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to read directory: %w", err)
	}

	// Initialize variables for counting and storing file full paths
	fileCount := 0
	var fileFullPaths []string

	// Loop through the entries
	for _, entry := range entries {
		if !entry.IsDir() {
			fileCount++
			fullPath := filepath.Join(dirPath, entry.Name()) // Construct the full file path
			fileFullPaths = append(fileFullPaths, fullPath)  // Add full path to the list
		}
	}

	return fileCount, fileFullPaths, nil
}
