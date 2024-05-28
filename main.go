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

	// Show the help page if no arguments are provided
	if len(os.Args) == 1 {
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
	var text string
	for scanner.Scan() {
		text += scanner.Text() + " "
	}
	text = strings.TrimSpace(text)

	// Translate the text
	translatedText, err := translateText(*fromLang, *toLang, text)
	if err != nil {
		fmt.Printf("Translation error: %v\n", err)
		os.Exit(1)
	}

	// Print the translated text
	fmt.Println(translatedText)
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

	// Translate the text
	translateResp, err := deepl.Translate(sourceLang, toLang, text)
	if err != nil {
		return "", err
	}

	// Extract the translated text from the deepl.Text type
	translatedText := translateResp.Result.Texts[0].Text

	// Return the translated text as a string
	return translatedText.Text, nil
}
