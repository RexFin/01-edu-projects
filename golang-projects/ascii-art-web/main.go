package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v\n", err)
		return
	}

	fmt.Fprintf(w, "POST Request successful\n")
	text := r.FormValue("text")
	bannerType := r.FormValue("banner-type")

	fmt.Fprintf(w, "Text = %s\n", text)
	fmt.Fprintf(w, "BannerType = %s\n", bannerType)

	fmt.Println(text)
	fmt.Println(bannerType)

}

func main() {
	http.HandleFunc("/ascii-art", formHandler)

	fileserver := http.FileServer(http.Dir("./template"))
	http.Handle("/", fileserver)

	fmt.Printf("Starting server at port 8181\n")

	serverError := http.ListenAndServe(":8181", nil)
	if serverError != nil {
		log.Fatal(serverError)
	}
}
