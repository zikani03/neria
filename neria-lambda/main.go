package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jdkato/prose/v2"
)

type NeriaEvent struct {
	URL      string `json:"url"`
	Selector string `json:"selector"`
}

type NERResult struct {
	Entities []NamedEntity `json:"entities"`
}

type NamedEntity struct {
	EntityType string `json:"entityType"`
	Name       string `json:"name"`
}

func NeriaEventHandler(ctx context.Context, event NeriaEvent) (NERResult, error) {
	dataString, err := scrapeData(event.URL, event.Selector)
	if err != nil {
		return NERResult{}, err
	}
	doc, _ := prose.NewDocument(dataString)
	result := NERResult{
		Entities: make([]NamedEntity, 0),
	}

	for _, ent := range doc.Entities() {
		result.Entities = append(result.Entities, NamedEntity{EntityType: ent.Label, Name: ent.Text})
	}

	return result, nil
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
	lambda.Start(NeriaEventHandler)
}
