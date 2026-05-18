package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Color struct {
	red   uint8
	green uint8
	blue  uint8
}

var baseColors = map[string]Color{
	"red":          {255, 0, 0},
	"orange":       {255, 125, 0},
	"yellow":       {255, 255, 0},
	"spring green": {125, 255, 0},
	"green":        {0, 255, 0},
	"turquoise":    {0, 255, 125},
	"cyan":         {0, 255, 255},
	"ocean":        {0, 125, 255},
	"blue":         {0, 0, 255},
	"violet":       {125, 0, 255},
	"magenta":      {255, 0, 255},
	"raspberry":    {255, 0, 125},
	"white":        {255, 255, 255},
	"black":        {0, 0, 0},
}

var expectedHashes = map[string]string{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "092d0cde973bfbb02522f18e00e8612e269f53bac358bb06f060a44abd0dbc52",
}

func checkIntegrity(filename string) (bool, error) {
	f, err := os.Open(filename)
	if err != nil {
		return false, err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return false, err
	}

	actualHash := fmt.Sprintf("%x", h.Sum(nil))
	expectedHash, ok := expectedHashes[filename]
	if !ok {
		return false, fmt.Errorf("unknown banner")
	}

	return actualHash == expectedHash, nil
}

func draw(word string, alphabet []string, startFillIndex, endFillIndex int, ColorRGB Color) {
	for i := 1; i <= 8; i++ {
		for j, char := range word {
			line := (int(char)-32)*9 + i

			if startFillIndex < 0 && endFillIndex < 0 {
				if line < len(alphabet) {
					fmt.Printf("\033[38;2;%d;%d;%dm%v\033[0m", ColorRGB.red, ColorRGB.green, ColorRGB.blue, alphabet[line])
				}
			} else {
				if line < len(alphabet) && j >= startFillIndex && j <= endFillIndex {
					fmt.Printf("\033[38;2;%d;%d;%dm%v\033[0m", ColorRGB.red, ColorRGB.green, ColorRGB.blue, alphabet[line])
				} else if line < len(alphabet) {
					fmt.Print(alphabet[line])
				}
			}

		}
		fmt.Println()
	}
}

func printExample() {
	fmt.Println()
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> \"something\"")
	fmt.Println()
}

func parseRGB(input string) (uint8, uint8, uint8) {
	re := regexp.MustCompile(`^rgb\(\s*(\d{1,3})\s*,\s*(\d{1,3})\s*,\s*(\d{1,3})\s*\)$`)

	matches := re.FindStringSubmatch(strings.ToLower(input))

	// If RGB input from user is incorrect
	// just return white color
	if matches == nil {
		return 255, 255, 255
	}

	r, errR := strconv.ParseUint(matches[1], 10, 8)
	g, errG := strconv.ParseUint(matches[2], 10, 8)
	b, errB := strconv.ParseUint(matches[3], 10, 8)

	// If RGB value not correct (out of range 0 to 255)
	// just return white color
	if errR != nil || errG != nil || errB != nil {
		return 255, 255, 255
	}

	return uint8(r), uint8(g), uint8(b)
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 || args[0] == "" || len(args) > 3 {
		printExample()
		os.Exit(1)
	}

	input := ""
	fillText := ""
	startFillIndex, endFillIndex := -1, -1
	color := args[0]
	colorRGB := Color{255, 255, 255} // Default white

	// For this project we only use standart banner
	banner := "standard.txt"

	if len(args) == 3 {
		fillText = args[1]
		input = args[2]
	} else if len(args) == 2 {
		input = args[1]
	}

	if !strings.Contains(color, "--color=") {
		printExample()
		os.Exit(1)
	}

	isValid, err := checkIntegrity(banner)
	if err != nil || !isValid {
		fmt.Printf("Error: banner file %s is corrupted\n", banner)
		os.Exit(1)
	}

	// Setuping color from baseColor
	colorName := strings.ToLower(string(strings.Split(color, "--color=")[1]))
	for k, v := range baseColors {
		if k == colorName {
			colorRGB = v
		}
	}

	// Setuping RGB (if user used it)
	if strings.Contains(colorName, "rgb") {
		colorRGB.red, colorRGB.green, colorRGB.blue = parseRGB(colorName)
	}

	// Finding range of filling text
	if fillText != "" {
		startFillIndex = strings.Index(input, fillText)
		if startFillIndex != -1 {
			endFillIndex = startFillIndex + (len(fillText) - 1)
		} else {
			// If text not found, just painting text to default
			colorRGB.red, colorRGB.green, colorRGB.blue = 255, 255, 255
		}
	}

	content, err := os.ReadFile(banner)
	if err != nil {
		fmt.Println("Error: could not find banner file", banner)
		return
	}

	alphabet := strings.Split(string(content), "\n")
	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}
		draw(word, alphabet, startFillIndex, endFillIndex, colorRGB)
	}
}
