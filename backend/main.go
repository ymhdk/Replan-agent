package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World!"))
	})

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe("8080", mux))
}
