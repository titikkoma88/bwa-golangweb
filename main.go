package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Welcome to Home"))
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello world, saya sedang belajar Golng Web"))
}