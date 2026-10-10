package main

import (
	"fmt"
	"go-refresh/internal/ai"
	"go-refresh/internal/modfields"
	"go-refresh/internal/parser"
	"go-refresh/internal/pipeline"
	"go-refresh/internal/wordcount"
	"log"
	"os"
	"strings"
)

// Ограничение в виде 5 мегабайтов
const maxFileSize = 5 * 1024 * 1024

func main() {

	if len(os.Args) < 3 || len(os.Args) > 4 {
		fmt.Println()
		fmt.Println("error: not valid arguments!")
		fmt.Println("usage: go run cmd/go-refresh/main.go <inputFile> <outputFile> <flag>")
		fmt.Println()
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fileInfo, err := os.Stat(inputFile)
	if err != nil {
		log.Fatalf("unable to get file info '%s': %v\n", inputFile, err)
	}
	if fileInfo.Size() > maxFileSize {
		log.Fatalf("error: file is too large! Maximum allowed size is 5MB.")
	}

	rawSlicesOfText, err := parser.ParseTextFromFile(inputFile)
	if err != nil {
		log.Fatalf("unable to open file '%s': %v\n", inputFile, err)
	}

	slicedText := modfields.ModificationFix([]rune(string(rawSlicesOfText)))
	slicedText = pipeline.Pipeline(slicedText)

	var result strings.Builder

	for i := 0; i < len(slicedText); i++ {
		word := slicedText[i]

		if word == "" {
			continue
		}

		if word == "\n" || word == "\r" {
			result.WriteByte('\n')
		} else {
			result.WriteString(word)

			if i+1 < len(slicedText) && slicedText[i+1] != "\n" && slicedText[i+1] != "\r" {
				result.WriteByte(' ')
			}
		}
	}

	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("unable to create file '%s': %v\n", outputFile, err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("error closing file: %v\n", err)
		}
	}()

	_, err = file.WriteString(result.String())
	if err != nil {
		log.Fatalf("unable to write to file '%s': %v\n", outputFile, err)
	}

	if len(os.Args) == 4 && os.Args[3] == "--lang" {
		ai.Langdetect(slicedText)
	}

	if len(os.Args) == 4 && os.Args[3] == "--keywords" {
		wordcount.Frequency(slicedText)
	}

	fmt.Printf("Success! Processed text saved to %s\n", outputFile)
}
