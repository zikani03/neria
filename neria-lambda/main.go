package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
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

func NeriaEventHandler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	var event NeriaEvent
	if err := json.Unmarshal([]byte(req.Body), &event); err != nil {
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	var dataString string

	if event.Text == "" {
		if event.URL == "" {
			return events.LambdaFunctionURLResponse{
				StatusCode: 400,
				Body:       fmt.Errorf("url cannot be empty").Error(),
			}, nil
		}

		if event.Selector == "" {
			return events.LambdaFunctionURLResponse{
				StatusCode: 400,
				Body:       fmt.Errorf("selector cannot be empty").Error(),
			}, nil
		}
		scrapedText, err := scrapeData(event.URL, event.Selector)
		if err != nil {
			return events.LambdaFunctionURLResponse{
				StatusCode: 500,
				Body:       err.Error(),
			}, nil
		}

		dataString = scrapedText

	} else {
		dataString = event.Text
	}

	if dataString == "" {
		return events.LambdaFunctionURLResponse{
			StatusCode: 400,
			Body:       fmt.Errorf("the text content is empty or could not be read").Error(),
		}, nil
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
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.LambdaFunctionURLResponse{Body: string(data), StatusCode: 200}, nil
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
