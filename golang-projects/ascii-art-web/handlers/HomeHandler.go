package handlers

import (
	"ascii-art-web/pkg"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorRender(w, &ErrorData{
			ErrorCode: 404,
			ErrorText: "This page not found",
		})
		return
	}

	Render(w, &PageData{
		Text:           "HELLO ASCII!",
		Banners:        pkg.GetBanners(),
		SelectedBanner: "standard",
		Result:         " _    _   ______   _        _         ____                      _____    _____   _____   _____   _  \n| |  | | |  ____| | |      | |       / __ \\            /\\      / ____|  / ____| |_   _| |_   _| | | \n| |__| | | |__    | |      | |      | |  | |          /  \\    | (___   | |        | |     | |   | | \n|  __  | |  __|   | |      | |      | |  | |         / /\\ \\    \\___ \\  | |        | |     | |   | | \n| |  | | | |____  | |____  | |____  | |__| |        / ____ \\   ____) | | |____   _| |_   _| |_  |_| \n|_|  |_| |______| |______| |______|  \\____/        /_/    \\_\\ |_____/   \\_____| |_____| |_____| (_) \n                                                                                                    \n                                                                                                    \n",
	})
}
