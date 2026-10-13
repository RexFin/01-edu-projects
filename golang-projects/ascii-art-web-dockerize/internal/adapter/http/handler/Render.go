package handler

import (
	"ascii-art-web/internal/domain"
	"log"
	"net/http"
	"text/template"
)

func Render(w http.ResponseWriter, data *domain.PageData) {
	index, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		log.Println(err)
		return
	}
	index.Execute(w, data)
}
