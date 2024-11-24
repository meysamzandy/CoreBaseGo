package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Specify the .env file
	filename := "configs/.env"

	// Call the function to update the .env file
	err := updateEnvFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(".env updated successfully")
	}
}

func updateEnvFile(filename string) error {
	// Read the .env file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var buffer bytes.Buffer
	scanner := bufio.NewScanner(file)
	var appKeyFound bool

	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains APP_KEY=
		if strings.HasPrefix(line, "APP_KEY=") {
			appKeyFound = true
			// Check if the APP_KEY value is empty
			if strings.TrimSpace(line) == "APP_KEY=" {
				// Generate a new key using OpenSSL
				newKey, err := generateOpenSSLKey()
				if err != nil {
					return err
				}
				line = fmt.Sprintf("APP_KEY=%s", newKey)
			}
		}

		// Append the line to the buffer
		buffer.WriteString(line + "\n")
	}

	// If APP_KEY was not found, append a new generated key
	if !appKeyFound {
		newKey, err := generateOpenSSLKey()
		if err != nil {
			return err
		}
		buffer.WriteString(fmt.Sprintf("APP_KEY=%s\n", newKey))
	}

	// Write the updated content back to the .env file
	err = os.WriteFile(filename, buffer.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

// generateOpenSSLKey generates a random base64 key using OpenSSL
func generateOpenSSLKey() (string, error) {
	cmd := exec.Command("openssl", "rand", "-base64", "32")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
