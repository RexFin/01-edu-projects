package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"

	httpDelivery "tour-track/http"
	"tour-track/repository"
)

//go:embed static templates/*.html
var webFiles embed.FS

func main() {
	templates, err := template.ParseFS(webFiles, "templates/*.html")
	if err != nil {
		log.Fatalf("load templates: %v", err)
	}
	mux := http.NewServeMux()
	staticContent, err := fs.Sub(webFiles, "static")
	if err != nil {
		log.Fatalf("load static files: %v", err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(staticContent))))
	httpDelivery.New(repository.NewClient(), templates).Register(mux)
	server := &http.Server{Addr: ":8080", Handler: mux, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 10 * time.Second, WriteTimeout: 20 * time.Second, IdleTimeout: 60 * time.Second}
	log.Printf("Tour Track is available at http://localhost%s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server stopped: %v", err)
	}
}
