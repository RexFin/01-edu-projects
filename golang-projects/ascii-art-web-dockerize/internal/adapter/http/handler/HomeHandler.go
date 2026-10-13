package handler

import (
	"ascii-art-web/internal/domain"
	"ascii-art-web/internal/service"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, &domain.PageData{
		Text:           "HELLO ASCII!",
		Banners:        service.GetBanners(),
		SelectedBanner: "standard",
		Result:         " _    _   ______   _        _         ____                      _____    _____   _____   _____   _  \n| |  | | |  ____| | |      | |       / __ \\            /\\      / ____|  / ____| |_   _| |_   _| | | \n| |__| | | |__    | |      | |      | |  | |          /  \\    | (___   | |        | |     | |   | | \n|  __  | |  __|   | |      | |      | |  | |         / /\\ \\    \\___ \\  | |        | |     | |   | | \n| |  | | | |____  | |____  | |____  | |__| |        / ____ \\   ____) | | |____   _| |_   _| |_  |_| \n|_|  |_| |______| |______| |______|  \\____/        /_/    \\_\\ |_____/   \\_____| |_____| |_____| (_) \n                                                                                                    \n                                                                                                    \n",
	})
}
