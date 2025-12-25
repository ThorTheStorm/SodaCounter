package fileops

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// Contants

// getUserHomeDir retrieves current working directory
func getUserHomeDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Could not find user profile path")
	}
	return userHomeDir
}

func GetInitPath() (string, error) {
	var jsonFilePath string
	userHomeDir := getUserHomeDir()
	//fmt.Printf("User home directory: %s\n", userHomeDir)
	if userHomeDir == "" {
		err := fmt.Errorf("Could not find user profile path")
		return "", err
	}
	if runtime.GOOS == "windows" {
		jsonFilePath = filepath.Join(userHomeDir, "\\Documents\\sodaCounter_data.json")
	} else {
		jsonFilePath = filepath.Join(userHomeDir, "//sodaCounter_data.json")
	}
	if jsonFilePath != "" {
		return jsonFilePath, nil
	} else {
		err := fmt.Errorf("Failed to get init-path")
		return "", err
	}
}

// ImportDataFromFile reads a JSON file and unmarshals it into a map[string]int.
func ImportDataFromFile(path string) (map[string]uint, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var result map[string]uint
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return result, nil
}

// ExportDataToFile marshals a map[string]int into JSON and writes it to a file.
func ExportDataToFile(path string, data map[string]uint) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
