package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

var templates *template.Template
var PORT = "8080"

func init() {
	templates = template.Must(template.ParseFiles("mainhtml.html"))
}

func main() {
	http.HandleFunc("/", homePageHandler)
	fmt.Println("Running server at http://localhost:" + PORT)
	fmt.Println("...to shut down server, press Ctrl+C")

	// files for the website
	http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

	http.ListenAndServe(":"+PORT, nil)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Bad Request: 404", http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodGet:
		// Get the text input from the query parameter
		text := r.URL.Query().Get("text")

		// Render the main page with the obtained text
		renderMainPage(w, text, nil)
	case http.MethodPost:
		// Parse form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad request: 404", http.StatusInternalServerError)
			return
		}

		text := r.FormValue("text")
		banner := r.FormValue("banner")

		// Check if a banner option is selected
		if banner == "" {
			http.Error(w, "Bad request: 400\nPlease select a banner type", http.StatusBadRequest)
			return
		}
		SymbolMap := banner + ".txt" // Select correct banner file

		// Generate ASCII art based on the selected banner
		asciiResult := AsciiArt(text, SymbolMap)
		// Render the result on the same page
		renderMainPage(w, text, asciiResult)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func renderMainPage(w http.ResponseWriter, text string, data interface{}) {
	// Create a struct to pass both text and data to the template
	type PageData struct {
		Text string
		Data interface{}
	}

	pageData := PageData{
		Text: text,
		Data: data,
	}

	err := templates.ExecuteTemplate(w, "mainhtml", pageData)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func AsciiArt(inputText string, SymbolMap string) string {
	var result strings.Builder

	m := make(map[rune][]string)
	content, err := os.ReadFile("templates/" + SymbolMap)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return result.String()
	}

	strContent := string(content)
	eachLine := strings.Split(strContent, "\n")

	for runes := ' '; '~' >= runes; runes++ {
		m[runes] = make([]string, 8)
		runeIndex := int(runes - ' ')
		for i := 0; i < 8; i++ {
			m[runes][i] = eachLine[runeIndex*9+i+1]
		}
	}

	lines := strings.Split(inputText, "\n")
	for _, line := range lines {
		if line == "" {
			result.WriteString("\n") // Preserve empty lines
		} else {
			for i := 0; i < 8; i++ {
				for _, lineCharacter := range line {
					if symbols, ok := m[lineCharacter]; ok {
						result.WriteString(symbols[i])
					} else {
						result.WriteString(" ") // Use a space for characters not in the map
					}
				}
				result.WriteString("\n") // Preserve line breaks between lines of ASCII art
			}
		}
	}
	return result.String()
}
