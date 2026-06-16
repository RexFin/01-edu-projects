package tests

import (
	"ascii-art-web/internal/service"
	"fmt"
	"os"
	"testing"
)

func TestGenerateAscii(t *testing.T) {
	originalDir, _ := os.Getwd()
	os.Chdir("..")
	defer os.Chdir(originalDir)

	var tests = []struct {
		name, text, banner string
		want               string
	}{
		{"Empty test", "", "", ""},
		{"Hello test", "Hello", "standard", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n"},
		{"Hello newline test", "Hello\n", "standard", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n\n"},
		{"Hello newline There test", "Hello\nThere", "standard", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n _______   _                           \n|__   __| | |                          \n   | |    | |__     ___   _ __    ___  \n   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n   | |    | | | | |  __/ | |    |  __/ \n   |_|    |_| |_|  \\___| |_|     \\___| \n                                       \n                                       \n"},
		{"Hello newline newline There test", "Hello\n\nThere", "standard", " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n\n _______   _                           \n|__   __| | |                          \n   | |    | |__     ___   _ __    ___  \n   | |    |  _ \\   / _ \\ | '__|  / _ \\ \n   | |    | | | | |  __/ | |    |  __/ \n   |_|    |_| |_|  \\___| |_|     \\___| \n                                       \n                                       \n"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s,%s", tt.text, tt.banner)
		t.Run(testname, func(t *testing.T) {
			ans, err := service.GenerateAscii(tt.text, tt.banner)
			if ans != tt.want {
				fmt.Println("GenerateAscii error:", err)
				t.Errorf("[%s] got \n%s, want \n%s", tt.name, ans, tt.want)
			}
		})
	}
}
