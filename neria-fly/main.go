package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/cors"
	"github.com/PuerkitoBio/goquery"
	"github.com/jdkato/prose/v2"
)

type NeriaEvent struct {
	URL      string `json:"Url"`
	Selector string `json:"Selector"`
	Text     string `json:"Text"`
}

type NERResult struct {
	Entities []NamedEntity `json:"Entities"`
}

type NamedEntity struct {
	EntityType string `json:"EntityType"`
	Name       string `json:"Name"`
}

func NeriaEventHandler(w http.ResponseWriter, r *http.Request) {
	var event NeriaEvent
	responseBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(responseBody, &event); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var dataString string

	if event.Text == "" {
		if event.URL == "" {
			http.Error(w, fmt.Errorf("url cannot be empty").Error(), http.StatusBadRequest)
			return
		}

		if event.Selector == "" {
			http.Error(w, fmt.Errorf("selector cannot be empty").Error(), http.StatusBadRequest)
			return
		}
		scrapedText, err := scrapeData(event.URL, event.Selector)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		dataString = scrapedText

	} else {
		dataString = event.Text
	}

	if dataString == "" {
		http.Error(w, fmt.Errorf("the text content is empty or could not be read").Error(), http.StatusBadRequest)
		return

	}

	doc, _ := prose.NewDocument(dataString)
	result := NERResult{
		Entities: make([]NamedEntity, 0),
	}

	for _, ent := range doc.Entities() {
		result.Entities = append(result.Entities, NamedEntity{EntityType: ent.Label, Name: ent.Text})
	}

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func scrapeData(url, domQuery string) (string, error) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to read data from URL: %v", err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return "", fmt.Errorf("failed to read data from URL: %v", err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to ready data fetched from URL: %v", err)
	}

	var sb strings.Builder
	// Find the review items
	doc.Find(domQuery).Each(func(i int, s *goquery.Selection) {
		// For each item found, get the text
		foundText := s.Text()
		sb.WriteString(foundText)
	})

	return sb.String(), nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	mux := http.NewServeMux()

	mux.HandleFunc("/", NeriaEventHandler)

	handler := cors.Default().Handler(mux)

	http.ListenAndServe(":"+port, handler)
}
