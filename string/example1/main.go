package main

import "fmt"

func main() {
    // Declare a string containing a Unicode character (😊)
    str := "😊"
    
    // Iterate over the string and print each character as a rune
    for _, char := range str {
        fmt.Printf("%c ", char)
    }
    fmt.Println()
}
