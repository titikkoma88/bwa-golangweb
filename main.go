package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/product", productHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request)  {
	log.Println(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to Home"))
}

func helloHandler(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("Hello world, saya sedang belajar Golng Web"))
}

func productHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.URL.Query().Get("id")
	idNumb, err := strconv.Atoi(id)
	
	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product Page : %d", idNumb)
}