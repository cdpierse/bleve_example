package main

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

// Article represents a story object from the HuffPost news category dataset.
type Article struct {
	Category         string `json:"category"`
	Headline         string `json:"headline"`
	Authors          string `json:"authors"`
	Link             string `json:"link"`
	ShortDescription string `json:"short_description"`
	Date             string `json:"date"`
}

// ReadDataset reads the news dataset provided by the string "name"
func ReadDataset(name string) ([]Article, error) {
	data, err := os.Open(name)
	if err != nil {
		log.Panic(err)
	}
	defer data.Close()

	var articles []Article

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {

		var tempArticle Article

		err := json.Unmarshal(scanner.Bytes(), &tempArticle)
		if err != nil {
			return nil, err
		}
		articles = append(articles, tempArticle)

	}
	return articles, nil

}
