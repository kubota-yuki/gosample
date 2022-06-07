package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, nil)
}

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/init", index)
	server.ListenAndServe()
}
