package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	fs := http.FileServer(http.Dir("images"))
	//http.HandleFunc("/", handler) // each request calls handler
	http.Handle("/images/", http.StripPrefix("/images/", fs))
	http.HandleFunc("/", serveTemplate)

	log.Println("Listening...")
	http.ListenAndServe(":8050", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path)
	lp := filepath.Join("index.html")
	tmpl, _ := template.ParseFiles(lp)
	tmpl.ExecuteTemplate(w, "index.html", "Hello")
}
