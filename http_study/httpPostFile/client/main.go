package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFileAndReturnResponse(filename string) string {
	fileDataBuffer := bytes.Buffer{}
	mpWritter := multipart.NewWriter(&fileDataBuffer)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	formFile, err := mpWritter.CreateFormFile("testFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(formFile, file); err != nil {
		log.Fatal(err)
	}
	mpWritter.Close()

	r, err := http.Post("http://localhost:8080", mpWritter.FormDataContentType(), &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := postFileAndReturnResponse("./testFile.txt")
	fmt.Println(data)
}
