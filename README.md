# Langa: A Simple Command-Line Text Translator
[![Go Report Card](https://goreportcard.com/badge/github.com/xplshn/langa)](https://goreportcard.com/report/github.com/xplshn/langa)
[![License](https://img.shields.io/badge/license-%20RABRMS-green)](https://github.com/xplshn/langa/blob/master/LICENSE)
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/xplshn/langa)

**Langa** is a minimalist command-line tool for translating text between languages without requiring an API key. It leverages the capabilities of the DeepL API to provide translations directly from your terminal, supporting a variety of languages.

## Features
- **No API Key Required**: Unlike many translation tools, Langa does not require an API key for basic usage, making it accessible for quick translations.
- **Unix Pipes**: Langa accepts input via pipes, allowing for integration into scripts or workflows.
- **Support for Multiple Languages**: Offers translation between numerous languages, with automatic detection of the source language if unspecified.

## (Non) Features
- **Non interactive interface**
- **Minimal, no new features will be added (probably)**

## Getting Started
Ensure you have Go installed on your system. Then, proceed to build and run Langa.

### Using Langa
To translate text, simply echo the text into Langa, specifying the target language with the `-to` flag. For example, to translate text to English:
```sh
echo "Bonjour le monde" | langa
```
Omit the `-from` flag to let Langa auto-detect the source language:
```sh
echo "Hello World" | langa -to fr
```

### Flags
- `-from`: Optional. Specifies the source language code (e.g., `es` for Spanish). Will default to auto-detect
- `-to`: Optional. Specifies the target language code (e.g., `en` for English). Will default to `en`

## License
Langa is licensed under the RABRMS License. This allows for the use, modification, and distribution of the software under certain conditions. For more details, please refer to the [LICENSE](LICENSE) file. This license is equivalent to the New or Revised BSD License.
