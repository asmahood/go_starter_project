package main

import (
	"fmt"
	"net/http"
)

/*
	-- Write a Function Challenge --

- Write a function that gets a URL and returns the value of the Content-Type response http header
- The function should perform an error if it cant perform a get request
*/

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	// make sure body is clsoed
	defer resp.Body.Close()

	ctype := resp.Header.Get("Content-Type")
	if ctype == "" {
		return "", fmt.Errorf("can't find content-type")
	}

	return ctype, nil
}

func main() {
	ctype, err := contentType("https://linkedin.com")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(ctype)
	}
}
