package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {

	fmt.Println("Hello, World!")
	PerformGetRequest()

}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"
	fmt.Println("Performing GET request to:", myUrl)
	response, err := http.Get(myUrl)
	fmt.Println("---------------")
	fmt.Println(response)
	fmt.Println("---------------")
	fmt.Println(err)
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

	var responseString strings.Builder
	content1, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content1)
	fmt.Println("Byte count:", byteCount)
	fmt.Println("Response body:", responseString.String())
}
