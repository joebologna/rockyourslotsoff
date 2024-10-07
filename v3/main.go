package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", serveHome)

	http.ListenAndServe("localhost:8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	html, _ := os.ReadFile("index.html")
	w.Write(html)
}
