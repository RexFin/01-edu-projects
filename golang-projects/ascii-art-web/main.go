package main

import (
	"ascii-art-web/handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/ascii-art", handlers.AsciiHandler)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil) // Maybe create own router...?
}
