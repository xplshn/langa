package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lib-x/deepl"
)

var (
	fromLang = flag.String("from", "", "Source language code (e.g., 'es' for Spanish)")
	toLang   = flag.String("to", "en", "Target language code (default is 'en' for English)")
)

func main() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()

	// Show the help page if no arguments are provided and no text is being piped
	if len(os.Args) == 1 && !isInputPiped() {
		flag.Usage()
		os.Exit(0)
	}

	// Check if the toLang flag is valid
	if !isValidLanguage(*toLang) {
		fmt.Println("Invalid target language code.")
		os.Exit(1)
	}

	// Read text input from the user
	scanner := bufio.NewScanner(os.Stdin)
	var text strings.Builder
	for scanner.Scan() {
		text.WriteString(scanner.Text() + " ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		os.Exit(1)
	}
	inputText := strings.TrimSpace(text.String())

	// Translate the text
	translatedText, err := translateText(*fromLang, *toLang, inputText)
	if err != nil {
		fmt.Printf("Translation error: %v\n", err)
		os.Exit(1)
	}

	// Print the translated text
	fmt.Println(translatedText)
}

// isInputPiped checks if the standard input is coming from a pipe
func isInputPiped() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

// isValidLanguage checks if a language code is valid
func isValidLanguage(langCode string) bool {
	// List of supported language codes
	supportedLangs := []string{"de", "en", "es", "fr", "it", "ja", "nl", "pl", "pt", "ru", "zh", "bg", "cs", "da", "el", "et", "fi", "hu", "lt", "lv", "ro", "sk", "sl", "sv"}

	// Check if the given language code exists in the supported list
	for _, lang := range supportedLangs {
		if langCode == lang {
			return true
		}
	}
	return false
}

// translateText translates text from source language to target language
func translateText(fromLang, toLang, text string) (string, error) {
	var sourceLang string
	if fromLang == "" {
		sourceLang = "auto" // Default source language is auto-detect
	} else {
		sourceLang = fromLang
	}

	// Translate the text using the deepl package
	translateResp, err := deepl.Translate(sourceLang, toLang, text)
	if err != nil {
		return "", err
	}

	// Extract the translated text from the deepl.Text type
	translatedText := translateResp.Result.Texts[0].Text

	return translatedText.Text, nil
}
