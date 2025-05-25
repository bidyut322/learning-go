package main

import (
	"fmt"
	"net/url"
)

func main() {
	myURL := "https://jsonplaceholder.typicode.com/posts/1"

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}
	fmt.Println("Scheme:", parsedURL.Scheme)
	fmt.Println("Host:", parsedURL.Host)
	fmt.Println("Path:", parsedURL.Path)
	fmt.Println("Raw Query:", parsedURL.RawQuery)

	// Modifying the URL
	parsedURL.Path = "/posts/2"
	parsedURL.RawQuery = "userId=1"

	newURL := parsedURL.String()
	fmt.Println("Modified URL:", newURL)
}
