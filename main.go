package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func result(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("result.html")
	t.Execute(w, nil)
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/result", result)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
