package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	fmt.Println("Hello, World!")
	PerformGetRequest()

}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Response status:", response.Status)
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Response body:", string(content))
	fmt.Println("Content length:", response.ContentLength)
	fmt.Println("Transfer encoding:", response.TransferEncoding)
	fmt.Println("Uncompressed:", response.Uncompressed)

	fmt.Println(content)
}
