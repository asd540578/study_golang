package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type test struct{}

func getDataWithCustomOptionsAndReturnResponse() string {
	req, err := http.NewRequest("GET", "http://localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "testToken")

	http.DefaultClient.Timeout = time.Second * 5

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := getDataWithCustomOptionsAndReturnResponse()
	fmt.Println(data)
}
