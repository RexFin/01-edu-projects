package app

import (
	"ascii-art-web/internal/adapter/http/server"
	"fmt"
	"log"
	"net/http"
)

func Run() {
	routes := server.NewRouter()

	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Println("Run error:", err)
	}
}
