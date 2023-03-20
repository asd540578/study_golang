package main

import (
	"html/template"
	"log"
	"net/http"
)

type Visitor struct {
	Name    string
	Surname string
	Age     string
}

type Hello struct {
	tmpl *template.Template
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	visitor := Visitor{}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	visitor.Name = r.Form.Get("name")
	visitor.Surname = r.Form.Get("surname")
	visitor.Age = r.Form.Get("age")

	h.tmpl.Execute(w, visitor)
}

func NewHello(tmplPath string) (*Hello, error) {
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return nil, err
	}

	return &Hello{tmpl}, nil
}

func main() {
	hello, err := NewHello("./result.html")
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/hello", hello)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./form.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
