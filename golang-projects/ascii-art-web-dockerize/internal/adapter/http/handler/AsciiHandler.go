package handler

import (
	"ascii-art-web/internal/domain"
	"ascii-art-web/internal/service"
	"ascii-art-web/pkg/validators"
	"net/http"
	"os"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorRender(w, &domain.ErrorData{
			ErrorCode: 405,
			ErrorText: "Method Not Allowed",
		})
		return
	}

	text := r.FormValue("text")
	selectedBanner := r.FormValue("banner")

	if len(text) > 10000 {
		ErrorRender(w, &domain.ErrorData{
			ErrorCode: 422,
			ErrorText: "Unprocessable Entity (symbols must be less than 10000)",
		})
		return
	}

	if text == "" || selectedBanner == "" || !validators.IsAscii(text) {
		ErrorRender(w, &domain.ErrorData{
			ErrorCode: 400,
			ErrorText: "Bad Request",
		})
		return
	}

	result, err := service.GenerateAscii(text, selectedBanner)
	if err != nil {
		if os.IsNotExist(err) {
			ErrorRender(w, &domain.ErrorData{
				ErrorCode: 404,
				ErrorText: "Banner Not Found",
			})
		} else {
			http.Error(w, "500 Internal Server Error", 500)
		}
		return
	}

	Render(w, &domain.PageData{
		Text:           text,
		Banners:        service.GetBanners(),
		SelectedBanner: selectedBanner,
		Result:         result,
	})
}
