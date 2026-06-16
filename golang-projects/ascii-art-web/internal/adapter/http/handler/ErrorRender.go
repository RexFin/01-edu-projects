package handler

import (
	"ascii-art-web/internal/domain"
	"html/template"
	"log"
	"net/http"
)

func ErrorRender(w http.ResponseWriter, errorData *domain.ErrorData) {
	index, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		log.Println(err)
		return
	}

	w.WriteHeader(errorData.ErrorCode)
	index.Execute(w, errorData)
}
