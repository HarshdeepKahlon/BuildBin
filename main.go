package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello from BuildBin</h1>"))
}

func showCodeSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a code snippet:"))
}

func createCodeSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Allow", "POST")
		http.Error(w, "Method is not Allowed", 405)
		return
	}
	w.Write([]byte("Create a new snippet:"))
}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showCodeSnippet)
	mux.HandleFunc("/snippet/create", createCodeSnippet)

	log.Println("Starting server on :4000")
	var err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
