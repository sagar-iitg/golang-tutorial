package main

import "fmt"

func main() {
	// String with escape sequences
	strWithEscape := "Hello\nWorld"

	// Convert the string to a byte slice using double quotes
	byteSliceWithQuotes := []byte(strWithEscape)
	fmt.Println("Byte slice with quotes:", byteSliceWithQuotes)

	// Convert the string to a byte slice using backticks
	byteSliceWithBackticks := []byte(`Hello\nWorld`)
	fmt.Println("Byte slice with backticks:", byteSliceWithBackticks)

	// Print each byte in the slices
	fmt.Println("Bytes in slice with quotes:")
	for _, b := range byteSliceWithQuotes {
		fmt.Printf("%c ", b)
	}
	fmt.Println("\nBytes in slice with backticks:")
	for _, b := range byteSliceWithBackticks {
		fmt.Printf("%c ", b)
	}
	fmt.Println()

	// Print the byte slices as strings
	fmt.Println("String from slice with quotes:", string(byteSliceWithQuotes))
	fmt.Println("String from slice with backticks:", string(byteSliceWithBackticks))
}
