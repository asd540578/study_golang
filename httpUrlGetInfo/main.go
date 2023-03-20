package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func hello(w http.ResponseWriter, r *http.Request) {
	v1 := r.URL.Query()
	name, ok := v1["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<h1>Missing name</h1>"))
		return
	}
	w.Write([]byte(fmt.Sprintf("<h1>Hello %s</h1>", strings.Join(name, ","))))
}

func main() {
	http.HandleFunc("/", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
