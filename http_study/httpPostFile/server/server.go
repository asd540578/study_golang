package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, fileHeader, err := r.FormFile("testFile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileContent, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fmt.Sprintf("./%s", fileHeader.Filename), fileContent, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s uploaded!", fileHeader.Filename)
	w.Write([]byte(fmt.Sprintf("%s uploaded!", fileHeader.Filename)))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
