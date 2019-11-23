package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from BuildBin"))
}

func main() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	var err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
