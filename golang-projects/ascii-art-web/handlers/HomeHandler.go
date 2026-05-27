package handlers

import (
	"ascii-art-web/pkg"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", 404)
		return
	}

	Render(w, &PageData{
		Text:    "Hello from HomeHandler!",
		Banners: pkg.GetBanners(),
	})
}
