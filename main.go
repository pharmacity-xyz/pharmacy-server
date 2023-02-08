package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	fmt.Printf("Starting the server on :%s... \n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
