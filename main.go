package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			name := "World"
			envName := os.Getenv("NAME")
			if envName != "" {
				name = envName
			}
			fmt.Fprintf(w, "Hello %s", name)
		})

	// Grab the PORT from the environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port :" + port)

	// Start the HTTP server on PORT
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
