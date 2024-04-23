package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"strings"
)

func calculateSHA256Hash(line string) string {
	// Calculate SHA256-Hash for each line
	hash := sha256.Sum256([]byte(line))
	return fmt.Sprintf("%x", hash)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// Remove CRLF and spaces
		line = strings.TrimSpace(line)
		// Calculate SHA256-Hash
		sha256Hash := calculateSHA256Hash(line)
		fmt.Println(sha256Hash)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
	}
}
