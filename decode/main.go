package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	// Your Base64 encoded string
	encodedString := "cVExTXNUN0RnclBQZkdYMllPeG9lRmtkSTJUaDdxbzE="

	// Decode the Base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Convert bytes to string (assuming it's a text string)
	decodedString := string(decodedBytes)

	fmt.Println(decodedString)
}
