package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Get user input for path, old character, and new character
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter path to scan: ")
	path, _ := reader.ReadString('\n')
	path = strings.TrimSpace(path)
	if path == "" {
		path = "." // Default to current directory
	}

	// Convert to absolute path
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("Error resolving absolute path:", err)
		return
	}

	fmt.Print("Enter character(s) to replace (use SPACE for spaces): ")
	oldChar, _ := reader.ReadString('\n')
	oldChar = strings.TrimSpace(oldChar)
	if oldChar == "SPACE" {
		oldChar = " "
	}
	if oldChar == "" {
		fmt.Println("You must specify a character to replace.")
		return
	}

	fmt.Print("Enter replacement character(s) (use SPACE for spaces): ")
	newChar, _ := reader.ReadString('\n')
	newChar = strings.TrimSpace(newChar)
	if newChar == "SPACE" {
		newChar = " "
	}

	// Read only the contents of the specified directory
	entries, err := os.ReadDir(absPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, entry := range entries {
		// Rename only if it's a directory and contains the oldChar
		if entry.IsDir() {
			dirName := entry.Name()
			if strings.Contains(dirName, oldChar) {
				newDirName := strings.ReplaceAll(dirName, oldChar, newChar)
				oldPath := filepath.Join(absPath, dirName)
				newPath := filepath.Join(absPath, newDirName)

				// Rename directory
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Printf("Error renaming folder '%s' to '%s': %v\n", oldPath, newPath, err)
				} else {
					fmt.Printf("Renamed '%s' to '%s'\n", oldPath, newPath)
				}
			}
		}
	}

	if err != nil {
		fmt.Println("Error scanning directory:", err)
	}
}
