package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080" // fallback
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go app on port %s!", port)
	})
	fmt.Println("Starting server on port", port)
	http.ListenAndServe(":"+port, nil)
}
