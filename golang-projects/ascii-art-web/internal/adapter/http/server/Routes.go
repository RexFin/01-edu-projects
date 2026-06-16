package server

import (
	"ascii-art-web/internal/adapter/http/handler"
	"ascii-art-web/internal/domain"
	"net/http"
	"strings"
)

type Routes struct{}

func NewRouter() *Routes {
	return &Routes{}
}

func (myRoutes Routes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/static/") {
		fs := http.FileServer(http.Dir("./static"))
		http.StripPrefix("/static/", fs).ServeHTTP(w, r)
		return
	}

	switch r.URL.Path {
	case "/":
		handler.HomeHandler(w, r)
	case "/ascii-art":
		handler.AsciiHandler(w, r)
	default:
		handler.ErrorRender(w, &domain.ErrorData{
			ErrorCode: 404,
			ErrorText: "Page Not Found",
		})
	}
}
