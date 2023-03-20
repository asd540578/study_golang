package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type server struct{}

type messageData struct {
	Message string `json:"message"`
}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := messageData{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(message)

	message.Message = strings.ToUpper(message.Message)
	jsonBytes, _ := json.Marshal(message)
	w.Write(jsonBytes)
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", server{}))
}
