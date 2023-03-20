package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func postDataAndReturnResponse(msg messageData) messageData {
	jsonBytes, _ := json.Marshal(msg)
	buffer := bytes.NewBuffer(jsonBytes)

	r, err := http.Post("http://localhost:8080", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	message := messageData{}
	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	msg := messageData{Message: "hi server!"}
	data := postDataAndReturnResponse(msg)

	fmt.Println(data.Message)
}
