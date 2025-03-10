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

	fmt.Print("Enter replacement character(s): ")
	newChar, _ := reader.ReadString('\n')
	newChar = strings.TrimSpace(newChar)

	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			replaceInFile(path, oldChar, newChar)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error scanning directory:", err)
	}
}

func replaceInFile(filePath, oldChar, newChar string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var content []string

	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, strings.ReplaceAll(line, oldChar, newChar))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	file.Close()
	file, err = os.Create(filePath)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range content {
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()
}
