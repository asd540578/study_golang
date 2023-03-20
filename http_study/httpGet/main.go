package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getDataAndReturnResponse() string {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := getDataAndReturnResponse()
	fmt.Print(data)
}
