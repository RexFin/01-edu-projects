package handlers

import (
	"ascii-art-web/pkg"
	"fmt"
	"net/http"
	"os"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "405 Method Not Allowed", 405)
		return
	}

	text := r.FormValue("text")
	selectedBanner := r.FormValue("banner")

	fmt.Printf("TEXT:\"%s\", BANNER:\"%s\"\n", text, selectedBanner)

	if text == "" || selectedBanner == "" || !pkg.IsAscii(text) {
		http.Error(w, "400 Bad Request", 400)
		fmt.Println("YA UPAL TUT")
		return
	}

	result, err := pkg.GenerateAscii(text, selectedBanner)
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "404 Banner Not Found", 404)
		} else {
			http.Error(w, "500 Internal Server Error", 500)
		}
		return
	}

	Render(w, &PageData{
		Text:           text,
		Banners:        pkg.GetBanners(),
		SelectedBanner: selectedBanner,
		Result:         result,
	})
}
