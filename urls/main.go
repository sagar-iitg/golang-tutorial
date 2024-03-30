package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://google.com"

func main() {

	fmt.Println("Hello Webrequest")

	//parsing the url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.RawPath)

}
