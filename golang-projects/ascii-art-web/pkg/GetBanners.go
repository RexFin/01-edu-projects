package pkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetBanners() []string {
	files, err := os.ReadDir("./banners/")
	if err != nil {
		fmt.Println("Error in GetBanners:", err)
		return []string{}
	}

	var banners []string

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()

		if filepath.Ext(fileName) == ".txt" {
			cleanName := strings.TrimSuffix(fileName, ".txt")
			banners = append(banners, cleanName)
		}
	}

	return banners
}
