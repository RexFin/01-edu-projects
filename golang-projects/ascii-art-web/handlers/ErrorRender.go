package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type ErrorData struct {
	ErrorCode int
	ErrorText string
}

func ErrorRender(w http.ResponseWriter, errorData *ErrorData) {
	index, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error", 500)
		log.Println(err)
		return
	}

	w.WriteHeader(errorData.ErrorCode)
	index.Execute(w, errorData)
}
