package handlers

import (
	"log"
	"net/http"
	"text/template"
)

type PageData struct {
	Text           string
	Banners        []string
	SelectedBanner string
	Result         string
}

func Render(w http.ResponseWriter, data *PageData) {
	index, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		log.Println(err)
		return
	}
	index.Execute(w, data)
}
